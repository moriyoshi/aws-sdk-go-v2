// Code generated by smithy-go-codegen DO NOT EDIT.

package forecast

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/forecast/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of dataset import jobs created using the CreateDatasetImportJob
// operation. For each import job, this operation returns a summary of its
// properties, including its Amazon Resource Name (ARN). You can retrieve the
// complete set of properties by using the ARN with the DescribeDatasetImportJob
// operation. You can filter the list by providing an array of Filter objects.
func (c *Client) ListDatasetImportJobs(ctx context.Context, params *ListDatasetImportJobsInput, optFns ...func(*Options)) (*ListDatasetImportJobsOutput, error) {
	if params == nil {
		params = &ListDatasetImportJobsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListDatasetImportJobs", params, optFns, addOperationListDatasetImportJobsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListDatasetImportJobsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListDatasetImportJobsInput struct {

	// An array of filters. For each filter, you provide a condition and a match
	// statement. The condition is either IS or IS_NOT, which specifies whether to
	// include or exclude the datasets that match the statement from the list,
	// respectively. The match statement consists of a key and a value. Filter
	// properties
	//
	//     * Condition - The condition to apply. Valid values are IS and
	// IS_NOT. To include the datasets that match the statement, specify IS. To exclude
	// matching datasets, specify IS_NOT.
	//
	//     * Key - The name of the parameter to
	// filter on. Valid values are DatasetArn and Status.
	//
	//     * Value - The value to
	// match.
	//
	// For example, to list all dataset import jobs whose status is ACTIVE, you
	// specify the following filter: "Filters": [ { "Condition": "IS", "Key": "Status",
	// "Value": "ACTIVE" } ]
	Filters []*types.Filter

	// The number of items to return in the response.
	MaxResults *int32

	// If the result of the previous request was truncated, the response includes a
	// NextToken. To retrieve the next set of results, use the token in the next
	// request. Tokens expire after 24 hours.
	NextToken *string
}

type ListDatasetImportJobsOutput struct {

	// An array of objects that summarize each dataset import job's properties.
	DatasetImportJobs []*types.DatasetImportJobSummary

	// If the response is truncated, Amazon Forecast returns this token. To retrieve
	// the next set of results, use the token in the next request.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListDatasetImportJobsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListDatasetImportJobs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListDatasetImportJobs{}, middleware.After)
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
	addOpListDatasetImportJobsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListDatasetImportJobs(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListDatasetImportJobs(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "forecast",
		OperationName: "ListDatasetImportJobs",
	}
}
