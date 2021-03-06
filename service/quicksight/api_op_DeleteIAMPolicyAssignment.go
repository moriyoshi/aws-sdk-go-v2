// Code generated by smithy-go-codegen DO NOT EDIT.

package quicksight

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes an existing IAM policy assignment.
func (c *Client) DeleteIAMPolicyAssignment(ctx context.Context, params *DeleteIAMPolicyAssignmentInput, optFns ...func(*Options)) (*DeleteIAMPolicyAssignmentOutput, error) {
	if params == nil {
		params = &DeleteIAMPolicyAssignmentInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteIAMPolicyAssignment", params, optFns, addOperationDeleteIAMPolicyAssignmentMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteIAMPolicyAssignmentOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteIAMPolicyAssignmentInput struct {

	// The name of the assignment.
	//
	// This member is required.
	AssignmentName *string

	// The AWS account ID where you want to delete the IAM policy assignment.
	//
	// This member is required.
	AwsAccountId *string

	// The namespace that contains the assignment.
	//
	// This member is required.
	Namespace *string
}

type DeleteIAMPolicyAssignmentOutput struct {

	// The name of the assignment.
	AssignmentName *string

	// The AWS request ID for this operation.
	RequestId *string

	// The HTTP status of the request.
	Status *int32

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteIAMPolicyAssignmentMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteIAMPolicyAssignment{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteIAMPolicyAssignment{}, middleware.After)
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
	addOpDeleteIAMPolicyAssignmentValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteIAMPolicyAssignment(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteIAMPolicyAssignment(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "quicksight",
		OperationName: "DeleteIAMPolicyAssignment",
	}
}
