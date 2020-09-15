// Code generated by smithy-go-codegen DO NOT EDIT.

package mediaconnect

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mediaconnect/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a new flow. The request must include one source. The request optionally
// can include outputs (up to 50) and entitlements (up to 50).
func (c *Client) CreateFlow(ctx context.Context, params *CreateFlowInput, optFns ...func(*Options)) (*CreateFlowOutput, error) {
	stack := middleware.NewStack("CreateFlow", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreateFlowMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpCreateFlowValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateFlow(options.Region), middleware.Before)

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, err
		}
	}
	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err := handler.Handle(ctx, params)
	if err != nil {
		return nil, &smithy.OperationError{
			ServiceID:     ServiceID,
			OperationName: "CreateFlow",
			Err:           err,
		}
	}
	out := result.(*CreateFlowOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Creates a new flow. The request must include one source. The request optionally
// can include outputs (up to 50) and entitlements (up to 50).
type CreateFlowInput struct {
	// The VPC interfaces you want on the flow.
	VpcInterfaces []*types.VpcInterfaceRequest
	// The Availability Zone that you want to create the flow in. These options are
	// limited to the Availability Zones within the current AWS Region.
	AvailabilityZone *string
	// The settings for source failover
	SourceFailoverConfig *types.FailoverConfig
	// The entitlements that you want to grant on a flow.
	Entitlements []*types.GrantEntitlementRequest
	Sources      []*types.SetSourceRequest
	// The outputs that you want to add to this flow.
	Outputs []*types.AddOutputRequest
	// The name of the flow.
	Name *string
	// The settings for the source of the flow.
	Source *types.SetSourceRequest
}

type CreateFlowOutput struct {
	// The settings for a flow, including its source, outputs, and entitlements.
	Flow *types.Flow

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreateFlowMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreateFlow{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateFlow{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateFlow(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mediaconnect",
		OperationName: "CreateFlow",
	}
}