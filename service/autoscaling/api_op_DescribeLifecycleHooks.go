// Code generated by smithy-go-codegen DO NOT EDIT.

package autoscaling

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes the lifecycle hooks for the specified Auto Scaling group.
func (c *Client) DescribeLifecycleHooks(ctx context.Context, params *DescribeLifecycleHooksInput, optFns ...func(*Options)) (*DescribeLifecycleHooksOutput, error) {
	if params == nil {
		params = &DescribeLifecycleHooksInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeLifecycleHooks", params, optFns, addOperationDescribeLifecycleHooksMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeLifecycleHooksOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeLifecycleHooksInput struct {

	// The name of the Auto Scaling group.
	//
	// This member is required.
	AutoScalingGroupName *string

	// The names of one or more lifecycle hooks. If you omit this parameter, all
	// lifecycle hooks are described.
	LifecycleHookNames []*string
}

type DescribeLifecycleHooksOutput struct {

	// The lifecycle hooks for the specified group.
	LifecycleHooks []*types.LifecycleHook

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeLifecycleHooksMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeLifecycleHooks{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeLifecycleHooks{}, middleware.After)
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
	addOpDescribeLifecycleHooksValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeLifecycleHooks(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeLifecycleHooks(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "autoscaling",
		OperationName: "DescribeLifecycleHooks",
	}
}
