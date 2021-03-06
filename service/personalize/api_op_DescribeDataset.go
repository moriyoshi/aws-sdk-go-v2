// Code generated by smithy-go-codegen DO NOT EDIT.

package personalize

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/personalize/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes the given dataset. For more information on datasets, see
// CreateDataset.
func (c *Client) DescribeDataset(ctx context.Context, params *DescribeDatasetInput, optFns ...func(*Options)) (*DescribeDatasetOutput, error) {
	if params == nil {
		params = &DescribeDatasetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeDataset", params, optFns, addOperationDescribeDatasetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeDatasetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeDatasetInput struct {

	// The Amazon Resource Name (ARN) of the dataset to describe.
	//
	// This member is required.
	DatasetArn *string
}

type DescribeDatasetOutput struct {

	// A listing of the dataset's properties.
	Dataset *types.Dataset

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeDatasetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeDataset{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeDataset{}, middleware.After)
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
	addOpDescribeDatasetValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeDataset(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeDataset(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "personalize",
		OperationName: "DescribeDataset",
	}
}
