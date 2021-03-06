// Code generated by smithy-go-codegen DO NOT EDIT.

package workspaces

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/workspaces/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Modifies the properties of the specified Amazon WorkSpaces clients.
func (c *Client) ModifyClientProperties(ctx context.Context, params *ModifyClientPropertiesInput, optFns ...func(*Options)) (*ModifyClientPropertiesOutput, error) {
	if params == nil {
		params = &ModifyClientPropertiesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyClientProperties", params, optFns, addOperationModifyClientPropertiesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyClientPropertiesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyClientPropertiesInput struct {

	// Information about the Amazon WorkSpaces client.
	//
	// This member is required.
	ClientProperties *types.ClientProperties

	// The resource identifiers, in the form of directory IDs.
	//
	// This member is required.
	ResourceId *string
}

type ModifyClientPropertiesOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationModifyClientPropertiesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpModifyClientProperties{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpModifyClientProperties{}, middleware.After)
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
	addOpModifyClientPropertiesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opModifyClientProperties(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opModifyClientProperties(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workspaces",
		OperationName: "ModifyClientProperties",
	}
}
