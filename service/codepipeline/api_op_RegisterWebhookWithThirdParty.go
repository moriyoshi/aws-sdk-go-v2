// Code generated by smithy-go-codegen DO NOT EDIT.

package codepipeline

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Configures a connection between the webhook that was created and the external
// tool with events to be detected.
func (c *Client) RegisterWebhookWithThirdParty(ctx context.Context, params *RegisterWebhookWithThirdPartyInput, optFns ...func(*Options)) (*RegisterWebhookWithThirdPartyOutput, error) {
	if params == nil {
		params = &RegisterWebhookWithThirdPartyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RegisterWebhookWithThirdParty", params, optFns, addOperationRegisterWebhookWithThirdPartyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RegisterWebhookWithThirdPartyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RegisterWebhookWithThirdPartyInput struct {

	// The name of an existing webhook created with PutWebhook to register with a
	// supported third party.
	WebhookName *string
}

type RegisterWebhookWithThirdPartyOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationRegisterWebhookWithThirdPartyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpRegisterWebhookWithThirdParty{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpRegisterWebhookWithThirdParty{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opRegisterWebhookWithThirdParty(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opRegisterWebhookWithThirdParty(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codepipeline",
		OperationName: "RegisterWebhookWithThirdParty",
	}
}
