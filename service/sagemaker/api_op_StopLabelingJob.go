// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Stops a running labeling job. A job that is stopped cannot be restarted. Any
// results obtained before the job is stopped are placed in the Amazon S3 output
// bucket.
func (c *Client) StopLabelingJob(ctx context.Context, params *StopLabelingJobInput, optFns ...func(*Options)) (*StopLabelingJobOutput, error) {
	if params == nil {
		params = &StopLabelingJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "StopLabelingJob", params, optFns, addOperationStopLabelingJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*StopLabelingJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type StopLabelingJobInput struct {

	// The name of the labeling job to stop.
	//
	// This member is required.
	LabelingJobName *string
}

type StopLabelingJobOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationStopLabelingJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpStopLabelingJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpStopLabelingJob{}, middleware.After)
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
	addOpStopLabelingJobValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opStopLabelingJob(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opStopLabelingJob(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "StopLabelingJob",
	}
}
