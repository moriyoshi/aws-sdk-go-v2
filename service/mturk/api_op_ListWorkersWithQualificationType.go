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

// The ListWorkersWithQualificationType operation returns all of the Workers that
// have been associated with a given Qualification type.
func (c *Client) ListWorkersWithQualificationType(ctx context.Context, params *ListWorkersWithQualificationTypeInput, optFns ...func(*Options)) (*ListWorkersWithQualificationTypeOutput, error) {
	if params == nil {
		params = &ListWorkersWithQualificationTypeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListWorkersWithQualificationType", params, optFns, addOperationListWorkersWithQualificationTypeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListWorkersWithQualificationTypeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListWorkersWithQualificationTypeInput struct {

	// The ID of the Qualification type of the Qualifications to return.
	//
	// This member is required.
	QualificationTypeId *string

	// Limit the number of results returned.
	MaxResults *int32

	// Pagination Token
	NextToken *string

	// The status of the Qualifications to return. Can be Granted | Revoked.
	Status types.QualificationStatus
}

type ListWorkersWithQualificationTypeOutput struct {

	// If the previous response was incomplete (because there is more data to
	// retrieve), Amazon Mechanical Turk returns a pagination token in the response.
	// You can use this pagination token to retrieve the next set of results.
	NextToken *string

	// The number of Qualifications on this page in the filtered results list,
	// equivalent to the number of Qualifications being returned by this call.
	NumResults *int32

	// The list of Qualification elements returned by this call.
	Qualifications []*types.Qualification

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListWorkersWithQualificationTypeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListWorkersWithQualificationType{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListWorkersWithQualificationType{}, middleware.After)
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
	addOpListWorkersWithQualificationTypeValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListWorkersWithQualificationType(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListWorkersWithQualificationType(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mturk-requester",
		OperationName: "ListWorkersWithQualificationType",
	}
}
