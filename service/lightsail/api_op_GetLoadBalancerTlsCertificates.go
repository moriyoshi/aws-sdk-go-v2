// Code generated by smithy-go-codegen DO NOT EDIT.

package lightsail

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lightsail/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns information about the TLS certificates that are associated with the
// specified Lightsail load balancer. TLS is just an updated, more secure version
// of Secure Socket Layer (SSL). You can have a maximum of 2 certificates
// associated with a Lightsail load balancer. One is active and the other is
// inactive.
func (c *Client) GetLoadBalancerTlsCertificates(ctx context.Context, params *GetLoadBalancerTlsCertificatesInput, optFns ...func(*Options)) (*GetLoadBalancerTlsCertificatesOutput, error) {
	stack := middleware.NewStack("GetLoadBalancerTlsCertificates", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpGetLoadBalancerTlsCertificatesMiddlewares(stack)
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
	addOpGetLoadBalancerTlsCertificatesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetLoadBalancerTlsCertificates(options.Region), middleware.Before)

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
			OperationName: "GetLoadBalancerTlsCertificates",
			Err:           err,
		}
	}
	out := result.(*GetLoadBalancerTlsCertificatesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetLoadBalancerTlsCertificatesInput struct {
	// The name of the load balancer you associated with your SSL/TLS certificate.
	LoadBalancerName *string
}

type GetLoadBalancerTlsCertificatesOutput struct {
	// An array of LoadBalancerTlsCertificate objects describing your SSL/TLS
	// certificates.
	TlsCertificates []*types.LoadBalancerTlsCertificate

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpGetLoadBalancerTlsCertificatesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpGetLoadBalancerTlsCertificates{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetLoadBalancerTlsCertificates{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetLoadBalancerTlsCertificates(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lightsail",
		OperationName: "GetLoadBalancerTlsCertificates",
	}
}