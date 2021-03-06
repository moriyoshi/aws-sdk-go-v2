// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes a logging level.
func (c *Client) DeleteV2LoggingLevel(ctx context.Context, params *DeleteV2LoggingLevelInput, optFns ...func(*Options)) (*DeleteV2LoggingLevelOutput, error) {
	if params == nil {
		params = &DeleteV2LoggingLevelInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteV2LoggingLevel", params, optFns, addOperationDeleteV2LoggingLevelMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteV2LoggingLevelOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteV2LoggingLevelInput struct {

	// The name of the resource for which you are configuring logging.
	//
	// This member is required.
	TargetName *string

	// The type of resource for which you are configuring logging. Must be THING_Group.
	//
	// This member is required.
	TargetType types.LogTargetType
}

type DeleteV2LoggingLevelOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteV2LoggingLevelMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteV2LoggingLevel{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteV2LoggingLevel{}, middleware.After)
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
	addOpDeleteV2LoggingLevelValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteV2LoggingLevel(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteV2LoggingLevel(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "DeleteV2LoggingLevel",
	}
}
