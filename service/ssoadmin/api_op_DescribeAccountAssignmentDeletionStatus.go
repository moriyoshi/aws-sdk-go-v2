// Code generated by smithy-go-codegen DO NOT EDIT.

package ssoadmin

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ssoadmin/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes the status of the assignment deletion request.
func (c *Client) DescribeAccountAssignmentDeletionStatus(ctx context.Context, params *DescribeAccountAssignmentDeletionStatusInput, optFns ...func(*Options)) (*DescribeAccountAssignmentDeletionStatusOutput, error) {
	if params == nil {
		params = &DescribeAccountAssignmentDeletionStatusInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeAccountAssignmentDeletionStatus", params, optFns, addOperationDescribeAccountAssignmentDeletionStatusMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeAccountAssignmentDeletionStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeAccountAssignmentDeletionStatusInput struct {

	// The identifier that is used to track the request operation progress.
	//
	// This member is required.
	AccountAssignmentDeletionRequestId *string

	// The ARN of the SSO instance under which the operation will be executed. For more
	// information about ARNs, see Amazon Resource Names (ARNs) and AWS Service
	// Namespaces in the AWS General Reference.
	//
	// This member is required.
	InstanceArn *string
}

type DescribeAccountAssignmentDeletionStatusOutput struct {

	// The status object for the account assignment deletion operation.
	AccountAssignmentDeletionStatus *types.AccountAssignmentOperationStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeAccountAssignmentDeletionStatusMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeAccountAssignmentDeletionStatus{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeAccountAssignmentDeletionStatus{}, middleware.After)
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
	addOpDescribeAccountAssignmentDeletionStatusValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeAccountAssignmentDeletionStatus(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeAccountAssignmentDeletionStatus(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sso",
		OperationName: "DescribeAccountAssignmentDeletionStatus",
	}
}
