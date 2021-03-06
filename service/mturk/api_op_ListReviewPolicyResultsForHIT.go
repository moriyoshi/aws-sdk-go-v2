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

// The ListReviewPolicyResultsForHIT operation retrieves the computed results and
// the actions taken in the course of executing your Review Policies for a given
// HIT. For information about how to specify Review Policies when you call
// CreateHIT, see Review Policies. The ListReviewPolicyResultsForHIT operation can
// return results for both Assignment-level and HIT-level review results.
func (c *Client) ListReviewPolicyResultsForHIT(ctx context.Context, params *ListReviewPolicyResultsForHITInput, optFns ...func(*Options)) (*ListReviewPolicyResultsForHITOutput, error) {
	if params == nil {
		params = &ListReviewPolicyResultsForHITInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListReviewPolicyResultsForHIT", params, optFns, addOperationListReviewPolicyResultsForHITMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListReviewPolicyResultsForHITOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListReviewPolicyResultsForHITInput struct {

	// The unique identifier of the HIT to retrieve review results for.
	//
	// This member is required.
	HITId *string

	// Limit the number of results returned.
	MaxResults *int32

	// Pagination token
	NextToken *string

	// The Policy Level(s) to retrieve review results for - HIT or Assignment. If
	// omitted, the default behavior is to retrieve all data for both policy levels.
	// For a list of all the described policies, see Review Policies.
	PolicyLevels []types.ReviewPolicyLevel

	// Specify if the operation should retrieve a list of the actions taken executing
	// the Review Policies and their outcomes.
	RetrieveActions *bool

	// Specify if the operation should retrieve a list of the results computed by the
	// Review Policies.
	RetrieveResults *bool
}

type ListReviewPolicyResultsForHITOutput struct {

	// The name of the Assignment-level Review Policy. This contains only the
	// PolicyName element.
	AssignmentReviewPolicy *types.ReviewPolicy

	// Contains both ReviewResult and ReviewAction elements for an Assignment.
	AssignmentReviewReport *types.ReviewReport

	// The HITId of the HIT for which results have been returned.
	HITId *string

	// The name of the HIT-level Review Policy. This contains only the PolicyName
	// element.
	HITReviewPolicy *types.ReviewPolicy

	// Contains both ReviewResult and ReviewAction elements for a particular HIT.
	HITReviewReport *types.ReviewReport

	// If the previous response was incomplete (because there is more data to
	// retrieve), Amazon Mechanical Turk returns a pagination token in the response.
	// You can use this pagination token to retrieve the next set of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListReviewPolicyResultsForHITMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListReviewPolicyResultsForHIT{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListReviewPolicyResultsForHIT{}, middleware.After)
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
	addOpListReviewPolicyResultsForHITValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListReviewPolicyResultsForHIT(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListReviewPolicyResultsForHIT(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mturk-requester",
		OperationName: "ListReviewPolicyResultsForHIT",
	}
}
