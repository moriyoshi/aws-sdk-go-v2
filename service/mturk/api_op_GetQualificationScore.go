// Code generated by smithy-go-codegen DO NOT EDIT.

package mturk

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mturk/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// The GetQualificationScore operation returns the value of a Worker's
// Qualification for a given Qualification type. To get a Worker's Qualification,
// you must know the Worker's ID. The Worker's ID is included in the assignment
// data returned by the ListAssignmentsForHIT operation. Only the owner of a
// Qualification type can query the value of a Worker's Qualification of that type.
func (c *Client) GetQualificationScore(ctx context.Context, params *GetQualificationScoreInput, optFns ...func(*Options)) (*GetQualificationScoreOutput, error) {
	if params == nil {
		params = &GetQualificationScoreInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetQualificationScore", params, optFns, addOperationGetQualificationScoreMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetQualificationScoreOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetQualificationScoreInput struct {

	// The ID of the QualificationType.
	//
	// This member is required.
	QualificationTypeId *string

	// The ID of the Worker whose Qualification is being updated.
	//
	// This member is required.
	WorkerId *string
}

type GetQualificationScoreOutput struct {

	// The Qualification data structure of the Qualification assigned to a user,
	// including the Qualification type and the value (score).
	Qualification *types.Qualification

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetQualificationScoreMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetQualificationScore{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetQualificationScore{}, middleware.After)
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
	addOpGetQualificationScoreValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetQualificationScore(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetQualificationScore(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mturk-requester",
		OperationName: "GetQualificationScore",
	}
}
