// Code generated by smithy-go-codegen DO NOT EDIT.

package batch

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes an AWS Batch compute environment. Before you can delete a compute
// environment, you must set its state to DISABLED with the
// UpdateComputeEnvironment () API operation and disassociate it from any job
// queues with the UpdateJobQueue () API operation.
func (c *Client) DeleteComputeEnvironment(ctx context.Context, params *DeleteComputeEnvironmentInput, optFns ...func(*Options)) (*DeleteComputeEnvironmentOutput, error) {
	stack := middleware.NewStack("DeleteComputeEnvironment", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpDeleteComputeEnvironmentMiddlewares(stack)
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
	addOpDeleteComputeEnvironmentValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteComputeEnvironment(options.Region), middleware.Before)

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
			OperationName: "DeleteComputeEnvironment",
			Err:           err,
		}
	}
	out := result.(*DeleteComputeEnvironmentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteComputeEnvironmentInput struct {
	// The name or Amazon Resource Name (ARN) of the compute environment to delete.
	ComputeEnvironment *string
}

type DeleteComputeEnvironmentOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpDeleteComputeEnvironmentMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpDeleteComputeEnvironment{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteComputeEnvironment{}, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteComputeEnvironment(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "batch",
		OperationName: "DeleteComputeEnvironment",
	}
}