// Code generated by smithy-go-codegen DO NOT EDIT.

package macie2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/macie2/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Declines Amazon Macie membership invitations that were received from specific
// accounts.
func (c *Client) DeclineInvitations(ctx context.Context, params *DeclineInvitationsInput, optFns ...func(*Options)) (*DeclineInvitationsOutput, error) {
	if params == nil {
		params = &DeclineInvitationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeclineInvitations", params, optFns, addOperationDeclineInvitationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeclineInvitationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeclineInvitationsInput struct {

	// An array that lists AWS account IDs, one for each account that sent an
	// invitation to decline.
	//
	// This member is required.
	AccountIds []*string
}

type DeclineInvitationsOutput struct {

	// An array of objects, one for each account whose invitation hasn't been declined.
	// Each object identifies the account and explains why the request hasn't been
	// processed for that account.
	UnprocessedAccounts []*types.UnprocessedAccount

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeclineInvitationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeclineInvitations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeclineInvitations{}, middleware.After)
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
	addOpDeclineInvitationsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeclineInvitations(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeclineInvitations(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "macie2",
		OperationName: "DeclineInvitations",
	}
}
