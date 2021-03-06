// Code generated by smithy-go-codegen DO NOT EDIT.

package schemas

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/schemas/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates the discoverer
func (c *Client) UpdateDiscoverer(ctx context.Context, params *UpdateDiscovererInput, optFns ...func(*Options)) (*UpdateDiscovererOutput, error) {
	if params == nil {
		params = &UpdateDiscovererInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateDiscoverer", params, optFns, addOperationUpdateDiscovererMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateDiscovererOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateDiscovererInput struct {

	// The ID of the discoverer.
	//
	// This member is required.
	DiscovererId *string

	// The description of the discoverer to update.
	Description *string
}

type UpdateDiscovererOutput struct {

	// The description of the discoverer.
	Description *string

	// The ARN of the discoverer.
	DiscovererArn *string

	// The ID of the discoverer.
	DiscovererId *string

	// The ARN of the event bus.
	SourceArn *string

	// The state of the discoverer.
	State types.DiscovererState

	// Tags associated with the resource.
	Tags map[string]*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateDiscovererMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateDiscoverer{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateDiscoverer{}, middleware.After)
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
	addOpUpdateDiscovererValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateDiscoverer(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opUpdateDiscoverer(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "schemas",
		OperationName: "UpdateDiscoverer",
	}
}
