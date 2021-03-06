// Code generated by smithy-go-codegen DO NOT EDIT.

package iotsitewise

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotsitewise/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves the current AWS IoT SiteWise logging options.
func (c *Client) DescribeLoggingOptions(ctx context.Context, params *DescribeLoggingOptionsInput, optFns ...func(*Options)) (*DescribeLoggingOptionsOutput, error) {
	if params == nil {
		params = &DescribeLoggingOptionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeLoggingOptions", params, optFns, addOperationDescribeLoggingOptionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeLoggingOptionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeLoggingOptionsInput struct {
}

type DescribeLoggingOptionsOutput struct {

	// The current logging options.
	//
	// This member is required.
	LoggingOptions *types.LoggingOptions

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeLoggingOptionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeLoggingOptions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeLoggingOptions{}, middleware.After)
	if err != nil {
		return err
	}
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	addResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	addRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addEndpointPrefix_opDescribeLoggingOptionsMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeLoggingOptions(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

type endpointPrefix_opDescribeLoggingOptionsMiddleware struct {
}

func (*endpointPrefix_opDescribeLoggingOptionsMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opDescribeLoggingOptionsMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.HostPrefix = "model."

	return next.HandleSerialize(ctx, in)
}
func addEndpointPrefix_opDescribeLoggingOptionsMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opDescribeLoggingOptionsMiddleware{}, `OperationSerializer`, middleware.Before)
}

func newServiceMetadataMiddleware_opDescribeLoggingOptions(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotsitewise",
		OperationName: "DescribeLoggingOptions",
	}
}
