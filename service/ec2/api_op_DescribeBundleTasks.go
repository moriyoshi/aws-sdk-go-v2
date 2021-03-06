// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes the specified bundle tasks or all of your bundle tasks. Completed
// bundle tasks are listed for only a limited time. If your bundle task is no
// longer in the list, you can still register an AMI from it. Just use
// RegisterImage with the Amazon S3 bucket name and image manifest name you
// provided to the bundle task.
func (c *Client) DescribeBundleTasks(ctx context.Context, params *DescribeBundleTasksInput, optFns ...func(*Options)) (*DescribeBundleTasksOutput, error) {
	if params == nil {
		params = &DescribeBundleTasksInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeBundleTasks", params, optFns, addOperationDescribeBundleTasksMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeBundleTasksOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeBundleTasksInput struct {

	// The bundle task IDs. Default: Describes all your bundle tasks.
	BundleIds []*string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool

	// The filters.
	//
	//     * bundle-id - The ID of the bundle task.
	//
	//     * error-code -
	// If the task failed, the error code returned.
	//
	//     * error-message - If the task
	// failed, the error message returned.
	//
	//     * instance-id - The ID of the
	// instance.
	//
	//     * progress - The level of task completion, as a percentage (for
	// example, 20%).
	//
	//     * s3-bucket - The Amazon S3 bucket to store the AMI.
	//
	//     *
	// s3-prefix - The beginning of the AMI name.
	//
	//     * start-time - The time the task
	// started (for example, 2013-09-15T17:15:20.000Z).
	//
	//     * state - The state of the
	// task (pending | waiting-for-shutdown | bundling | storing | cancelling |
	// complete | failed).
	//
	//     * update-time - The time of the most recent update for
	// the task.
	Filters []*types.Filter
}

type DescribeBundleTasksOutput struct {

	// Information about the bundle tasks.
	BundleTasks []*types.BundleTask

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeBundleTasksMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDescribeBundleTasks{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDescribeBundleTasks{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeBundleTasks(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeBundleTasks(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DescribeBundleTasks",
	}
}
