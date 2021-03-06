// Code generated by smithy-go-codegen DO NOT EDIT.

package macie2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/macie2/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Retrieves information about a member account that's associated with an Amazon
// Macie master account.
func (c *Client) GetMember(ctx context.Context, params *GetMemberInput, optFns ...func(*Options)) (*GetMemberOutput, error) {
	if params == nil {
		params = &GetMemberInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetMember", params, optFns, addOperationGetMemberMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetMemberOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetMemberInput struct {

	// The unique identifier for the Amazon Macie resource or account that the request
	// applies to.
	//
	// This member is required.
	Id *string
}

type GetMemberOutput struct {

	// The AWS account ID for the account.
	AccountId *string

	// The Amazon Resource Name (ARN) of the account.
	Arn *string

	// The email address for the account.
	Email *string

	// The date and time, in UTC and extended ISO 8601 format, when an Amazon Macie
	// membership invitation was last sent to the account. This value is null if a
	// Macie invitation hasn't been sent to the account.
	InvitedAt *time.Time

	// The AWS account ID for the master account.
	MasterAccountId *string

	// The current status of the relationship between the account and the master
	// account.
	RelationshipStatus types.RelationshipStatus

	// A map of key-value pairs that identifies the tags (keys and values) that are
	// associated with the member account in Amazon Macie.
	Tags map[string]*string

	// The date and time, in UTC and extended ISO 8601 format, of the most recent
	// change to the status of the relationship between the account and the master
	// account.
	UpdatedAt *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetMemberMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetMember{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetMember{}, middleware.After)
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
	addOpGetMemberValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetMember(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetMember(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "macie2",
		OperationName: "GetMember",
	}
}
