// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves a sortable, filterable list of existing AWS Glue machine learning
// transforms in this AWS account, or the resources with the specified tag. This
// operation takes the optional Tags field, which you can use as a filter of the
// responses so that tagged resources can be retrieved as a group. If you choose to
// use tag filtering, only resources with the tags are retrieved.
func (c *Client) ListMLTransforms(ctx context.Context, params *ListMLTransformsInput, optFns ...func(*Options)) (*ListMLTransformsOutput, error) {
	if params == nil {
		params = &ListMLTransformsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListMLTransforms", params, optFns, addOperationListMLTransformsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListMLTransformsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListMLTransformsInput struct {

	// A TransformFilterCriteria used to filter the machine learning transforms.
	Filter *types.TransformFilterCriteria

	// The maximum size of a list to return.
	MaxResults *int32

	// A continuation token, if this is a continuation request.
	NextToken *string

	// A TransformSortCriteria used to sort the machine learning transforms.
	Sort *types.TransformSortCriteria

	// Specifies to return only these tagged resources.
	Tags map[string]*string
}

type ListMLTransformsOutput struct {

	// The identifiers of all the machine learning transforms in the account, or the
	// machine learning transforms with the specified tags.
	//
	// This member is required.
	TransformIds []*string

	// A continuation token, if the returned list does not contain the last metric
	// available.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListMLTransformsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListMLTransforms{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListMLTransforms{}, middleware.After)
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
	addOpListMLTransformsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListMLTransforms(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListMLTransforms(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "ListMLTransforms",
	}
}
