package customizations

import (
	"context"
	"fmt"
	"strings"

	"github.com/awslabs/smithy-go/middleware"
	"github.com/awslabs/smithy-go/transport/http"

	awsarn "github.com/aws/aws-sdk-go-v2/aws/arn"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/internal/s3shared"
	"github.com/aws/aws-sdk-go-v2/service/internal/s3shared/arn"
)

// processARNResourceMiddleware is used to process an ARN resource.
type processARNResourceMiddleware struct {
	// GetARN points to a function that processes an input and returns ARN,
	// and bool indicating if ARN is supported or set.
	GetARN func(interface{}) (string, bool)

	// UseARNRegion indicates if region parsed from an ARN should be used.
	UseARNRegion bool
}

// ID returns the middleware ID.
func (*processARNResourceMiddleware) ID() string { return "S3:ProcessARNResourceMiddleware" }

func (m *processARNResourceMiddleware) HandleSerialize(
	ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler,
) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	req, ok := in.Request.(*http.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown request type %T", req)
	}

	// get arn from an arnable field.
	v, ok := m.GetARN(in.Parameters)
	if !ok {
		return next.HandleSerialize(ctx, in)
	}

	resource, err := parseEndpointARN(v)
	if err != nil {
		return out, metadata, fmt.Errorf("Error parsing resource arn %w", err)
	}

	resourceRequest := s3shared.ResourceRequest{
		Resource: resource,
		Request:  req,
		Options: s3shared.ResourceRequestOptions{
			RequestRegion: awsmiddleware.GetRegion(ctx),
			SigningRegion: awsmiddleware.GetSigningRegion(ctx),
			PartitionID:   awsmiddleware.GetPartitionID(ctx),
			UseARNRegion:  m.UseARNRegion,
			// TODO: change HasCustomEndpoint value to use HostImmutable from ctx
			HasCustomEndpoint: false,
		},
	}

	// validate resource request
	if err := validateResourceRequest(resourceRequest); err != nil {
		return out, metadata, err
	}

	// switch to

	return next.HandleSerialize(ctx, in)
}

// validate if s3 resource and request config is compatible.
func validateResourceRequest(resourceRequest s3shared.ResourceRequest) error {
	// check if resourceRequest leads to a cross partition error
	v, err := resourceRequest.IsCrossPartition()
	if err != nil {
		return err
	}
	if v {
		// if cross partition
		return s3shared.NewClientPartitionMismatchError(resourceRequest.Resource,
			resourceRequest.Options.PartitionID, resourceRequest.Options.RequestRegion, nil)
	}

	// check if resourceRequest leads to a cross region error
	v, err = resourceRequest.IsCrossRegion()
	if err != nil {
		return err
	}
	if !resourceRequest.AllowCrossRegion() && v {
		// if cross region, but not use ARN region is not enabled
		return s3shared.NewClientRegionMismatchError(resourceRequest.Resource,
			resourceRequest.Options.PartitionID, resourceRequest.Options.RequestRegion, nil)
	}

	// check if resourceRequest leads to a custom endpoint error
	if resourceRequest.HasCustomEndpoint() {
		// if a custom endpoint was provided, return an error
		return s3shared.NewInvalidARNWithCustomEndpointError(resourceRequest.Resource, nil)
	}

	return nil
}

// Used by shapes with members decorated as endpoint ARN.
func parseEndpointARN(v string) (arn.Resource, error) {
	return arn.ParseResource(v, accessPointResourceParser)
}

func accessPointResourceParser(a awsarn.ARN) (arn.Resource, error) {
	resParts := arn.SplitResource(a.Resource)
	switch resParts[0] {
	case "accesspoint":
		if a.Service != "s3" {
			return arn.AccessPointARN{}, arn.InvalidARNError{ARN: a, Reason: "service is not s3"}
		}
		return arn.ParseAccessPointResource(a, resParts[1:])
	case "outpost":
		if a.Service != "s3-outposts" {
			return arn.OutpostAccessPointARN{}, arn.InvalidARNError{ARN: a, Reason: "service is not s3-outposts"}
		}
		return parseOutpostAccessPointResource(a, resParts[1:])
	default:
		return nil, arn.InvalidARNError{ARN: a, Reason: "unknown resource type"}
	}
}

func parseOutpostAccessPointResource(a awsarn.ARN, resParts []string) (arn.OutpostAccessPointARN, error) {
	// outpost accesspoint arn is only valid if service is s3-outposts
	if a.Service != "s3-outposts" {
		return arn.OutpostAccessPointARN{}, arn.InvalidARNError{ARN: a, Reason: "service is not s3-outposts"}
	}

	if len(resParts) == 0 {
		return arn.OutpostAccessPointARN{}, arn.InvalidARNError{ARN: a, Reason: "outpost resource-id not set"}
	}

	if len(resParts) < 3 {
		return arn.OutpostAccessPointARN{}, arn.InvalidARNError{
			ARN: a, Reason: "access-point resource not set in Outpost ARN",
		}
	}

	resID := strings.TrimSpace(resParts[0])
	if len(resID) == 0 {
		return arn.OutpostAccessPointARN{}, arn.InvalidARNError{ARN: a, Reason: "outpost resource-id not set"}
	}

	var outpostAccessPointARN = arn.OutpostAccessPointARN{}
	switch resParts[1] {
	case "accesspoint":
		accessPointARN, err := arn.ParseAccessPointResource(a, resParts[2:])
		if err != nil {
			return arn.OutpostAccessPointARN{}, err
		}
		// set access-point arn
		outpostAccessPointARN.AccessPointARN = accessPointARN
	default:
		return arn.OutpostAccessPointARN{}, arn.InvalidARNError{ARN: a, Reason: "access-point resource not set in Outpost ARN"}
	}

	// set outpost id
	outpostAccessPointARN.OutpostID = resID
	return outpostAccessPointARN, nil
}
