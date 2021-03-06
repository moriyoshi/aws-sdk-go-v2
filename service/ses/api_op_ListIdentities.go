// Code generated by smithy-go-codegen DO NOT EDIT.

package ses

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ses/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list containing all of the identities (email addresses and domains)
// for your AWS account in the current AWS Region, regardless of verification
// status. You can execute this operation no more than once per second.
func (c *Client) ListIdentities(ctx context.Context, params *ListIdentitiesInput, optFns ...func(*Options)) (*ListIdentitiesOutput, error) {
	if params == nil {
		params = &ListIdentitiesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListIdentities", params, optFns, addOperationListIdentitiesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListIdentitiesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents a request to return a list of all identities (email addresses and
// domains) that you have attempted to verify under your AWS account, regardless of
// verification status.
type ListIdentitiesInput struct {

	// The type of the identities to list. Possible values are "EmailAddress" and
	// "Domain". If this parameter is omitted, then all identities will be listed.
	IdentityType types.IdentityType

	// The maximum number of identities per page. Possible values are 1-1000 inclusive.
	MaxItems *int32

	// The token to use for pagination.
	NextToken *string
}

// A list of all identities that you have attempted to verify under your AWS
// account, regardless of verification status.
type ListIdentitiesOutput struct {

	// A list of identities.
	//
	// This member is required.
	Identities []*string

	// The token used for pagination.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListIdentitiesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpListIdentities{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpListIdentities{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListIdentities(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListIdentities(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ses",
		OperationName: "ListIdentities",
	}
}
