// Code generated by smithy-go-codegen DO NOT EDIT.

package medialive

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/medialive/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Stops running resources
func (c *Client) BatchStop(ctx context.Context, params *BatchStopInput, optFns ...func(*Options)) (*BatchStopOutput, error) {
	if params == nil {
		params = &BatchStopInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchStop", params, optFns, addOperationBatchStopMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchStopOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to stop resources
type BatchStopInput struct {

	// List of channel IDs
	ChannelIds []*string

	// List of multiplex IDs
	MultiplexIds []*string
}

// Placeholder documentation for BatchStopResponse
type BatchStopOutput struct {

	// List of failed operations
	Failed []*types.BatchFailedResultModel

	// List of successful operations
	Successful []*types.BatchSuccessfulResultModel

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationBatchStopMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpBatchStop{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpBatchStop{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opBatchStop(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opBatchStop(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "medialive",
		OperationName: "BatchStop",
	}
}
