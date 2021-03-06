// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets the specified identity provider.
func (c *Client) GetIdentityProviderByIdentifier(ctx context.Context, params *GetIdentityProviderByIdentifierInput, optFns ...func(*Options)) (*GetIdentityProviderByIdentifierOutput, error) {
	if params == nil {
		params = &GetIdentityProviderByIdentifierInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetIdentityProviderByIdentifier", params, optFns, addOperationGetIdentityProviderByIdentifierMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetIdentityProviderByIdentifierOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetIdentityProviderByIdentifierInput struct {

	// The identity provider ID.
	//
	// This member is required.
	IdpIdentifier *string

	// The user pool ID.
	//
	// This member is required.
	UserPoolId *string
}

type GetIdentityProviderByIdentifierOutput struct {

	// The identity provider object.
	//
	// This member is required.
	IdentityProvider *types.IdentityProviderType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetIdentityProviderByIdentifierMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetIdentityProviderByIdentifier{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetIdentityProviderByIdentifier{}, middleware.After)
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
	addOpGetIdentityProviderByIdentifierValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetIdentityProviderByIdentifier(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetIdentityProviderByIdentifier(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-idp",
		OperationName: "GetIdentityProviderByIdentifier",
	}
}
