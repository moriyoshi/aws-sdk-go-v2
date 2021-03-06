// Code generated by smithy-go-codegen DO NOT EDIT.

package datapipeline

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/datapipeline/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves metadata about one or more pipelines. The information retrieved
// includes the name of the pipeline, the pipeline identifier, its current state,
// and the user account that owns the pipeline. Using account credentials, you can
// retrieve metadata about pipelines that you or your IAM users have created. If
// you are using an IAM user account, you can retrieve metadata about only those
// pipelines for which you have read permissions. To retrieve the full pipeline
// definition instead of metadata about the pipeline, call GetPipelineDefinition.
func (c *Client) DescribePipelines(ctx context.Context, params *DescribePipelinesInput, optFns ...func(*Options)) (*DescribePipelinesOutput, error) {
	if params == nil {
		params = &DescribePipelinesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribePipelines", params, optFns, addOperationDescribePipelinesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribePipelinesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Contains the parameters for DescribePipelines.
type DescribePipelinesInput struct {

	// The IDs of the pipelines to describe. You can pass as many as 25 identifiers in
	// a single call. To obtain pipeline IDs, call ListPipelines.
	//
	// This member is required.
	PipelineIds []*string
}

// Contains the output of DescribePipelines.
type DescribePipelinesOutput struct {

	// An array of descriptions for the specified pipelines.
	//
	// This member is required.
	PipelineDescriptionList []*types.PipelineDescription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribePipelinesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribePipelines{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribePipelines{}, middleware.After)
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
	addOpDescribePipelinesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribePipelines(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribePipelines(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "datapipeline",
		OperationName: "DescribePipelines",
	}
}
