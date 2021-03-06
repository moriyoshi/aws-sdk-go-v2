// Code generated by smithy-go-codegen DO NOT EDIT.

package securityhub

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/securityhub/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns a list of the available standards in Security Hub. For each standard,
// the results include the standard ARN, the name, and a description.
func (c *Client) DescribeStandards(ctx context.Context, params *DescribeStandardsInput, optFns ...func(*Options)) (*DescribeStandardsOutput, error) {
	if params == nil {
		params = &DescribeStandardsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeStandards", params, optFns, addOperationDescribeStandardsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeStandardsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeStandardsInput struct {

	// The maximum number of standards to return.
	MaxResults *int32

	// The token that is required for pagination. On your first call to the
	// DescribeStandards operation, set the value of this parameter to NULL. For
	// subsequent calls to the operation, to continue listing data, set the value of
	// this parameter to the value returned from the previous response.
	NextToken *string
}

type DescribeStandardsOutput struct {

	// The pagination token to use to request the next page of results.
	NextToken *string

	// A list of available standards.
	Standards []*types.Standard

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeStandardsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeStandards{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeStandards{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeStandards(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeStandards(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "securityhub",
		OperationName: "DescribeStandards",
	}
}
