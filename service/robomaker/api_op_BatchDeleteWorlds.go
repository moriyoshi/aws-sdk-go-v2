// Code generated by smithy-go-codegen DO NOT EDIT.

package robomaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes one or more worlds in a batch operation.
func (c *Client) BatchDeleteWorlds(ctx context.Context, params *BatchDeleteWorldsInput, optFns ...func(*Options)) (*BatchDeleteWorldsOutput, error) {
	if params == nil {
		params = &BatchDeleteWorldsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchDeleteWorlds", params, optFns, addOperationBatchDeleteWorldsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchDeleteWorldsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchDeleteWorldsInput struct {

	// A list of Amazon Resource Names (arns) that correspond to worlds to delete.
	//
	// This member is required.
	Worlds []*string
}

type BatchDeleteWorldsOutput struct {

	// A list of unprocessed worlds associated with the call. These worlds were not
	// deleted.
	UnprocessedWorlds []*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationBatchDeleteWorldsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpBatchDeleteWorlds{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpBatchDeleteWorlds{}, middleware.After)
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
	addOpBatchDeleteWorldsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opBatchDeleteWorlds(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opBatchDeleteWorlds(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "robomaker",
		OperationName: "BatchDeleteWorlds",
	}
}
