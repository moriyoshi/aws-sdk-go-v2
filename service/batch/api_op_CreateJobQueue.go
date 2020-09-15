// Code generated by smithy-go-codegen DO NOT EDIT.

package batch

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/batch/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates an AWS Batch job queue. When you create a job queue, you associate one
// or more compute environments to the queue and assign an order of preference for
// the compute environments. You also set a priority to the job queue that
// determines the order in which the AWS Batch scheduler places jobs onto its
// associated compute environments. For example, if a compute environment is
// associated with more than one job queue, the job queue with a higher priority is
// given preference for scheduling jobs to that compute environment.
func (c *Client) CreateJobQueue(ctx context.Context, params *CreateJobQueueInput, optFns ...func(*Options)) (*CreateJobQueueOutput, error) {
	stack := middleware.NewStack("CreateJobQueue", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreateJobQueueMiddlewares(stack)
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
	addOpCreateJobQueueValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateJobQueue(options.Region), middleware.Before)

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
			OperationName: "CreateJobQueue",
			Err:           err,
		}
	}
	out := result.(*CreateJobQueueOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateJobQueueInput struct {
	// The name of the job queue.
	JobQueueName *string
	// The set of compute environments mapped to a job queue and their order relative
	// to each other. The job scheduler uses this parameter to determine which compute
	// environment should execute a given job. Compute environments must be in the
	// VALID state before you can associate them with a job queue. You can associate up
	// to three compute environments with a job queue.
	ComputeEnvironmentOrder []*types.ComputeEnvironmentOrder
	// The state of the job queue. If the job queue state is ENABLED, it is able to
	// accept jobs.
	State types.JQState
	// The priority of the job queue. Job queues with a higher priority (or a higher
	// integer value for the priority parameter) are evaluated first when associated
	// with the same compute environment. Priority is determined in descending order,
	// for example, a job queue with a priority value of 10 is given scheduling
	// preference over a job queue with a priority value of 1.
	Priority *int32
}

type CreateJobQueueOutput struct {
	// The name of the job queue.
	JobQueueName *string
	// The Amazon Resource Name (ARN) of the job queue.
	JobQueueArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreateJobQueueMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreateJobQueue{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateJobQueue{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateJobQueue(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "batch",
		OperationName: "CreateJobQueue",
	}
}