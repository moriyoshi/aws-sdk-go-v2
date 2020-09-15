// Code generated by smithy-go-codegen DO NOT EDIT.

package ecr

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ecr/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Lists all the image IDs for the specified repository. You can filter images
// based on whether or not they are tagged by using the tagStatus filter and
// specifying either TAGGED, UNTAGGED or ANY. For example, you can filter your
// results to return only UNTAGGED images and then pipe that result to a
// BatchDeleteImage () operation to delete them. Or, you can filter your results to
// return only TAGGED images to list all of the tags in your repository.
func (c *Client) ListImages(ctx context.Context, params *ListImagesInput, optFns ...func(*Options)) (*ListImagesOutput, error) {
	stack := middleware.NewStack("ListImages", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpListImagesMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	addHTTPSignerV4Middleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	addClientUserAgent(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpListImagesValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListImages(options.Region), middleware.Before)

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, err
		}
	}
	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err := handler.Handle(ctx, params)
	if err != nil {
		return nil, &smithy.OperationError{
			ServiceID:     ServiceID,
			OperationName: "ListImages",
			Err:           err,
		}
	}
	out := result.(*ListImagesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListImagesInput struct {
	// The nextToken value returned from a previous paginated ListImages request where
	// maxResults was used and the results exceeded the value of that parameter.
	// Pagination continues from the end of the previous results that returned the
	// nextToken value. This value is null when there are no more results to return.
	// This token should be treated as an opaque identifier that is only used to
	// retrieve the next items in a list and not for other programmatic purposes.
	NextToken *string
	// The maximum number of image results returned by ListImages in paginated output.
	// When this parameter is used, ListImages only returns maxResults results in a
	// single page along with a nextToken response element. The remaining results of
	// the initial request can be seen by sending another ListImages request with the
	// returned nextToken value. This value can be between 1 and 1000. If this
	// parameter is not used, then ListImages returns up to 100 results and a nextToken
	// value, if applicable.
	MaxResults *int32
	// The filter key and value with which to filter your ListImages results.
	Filter *types.ListImagesFilter
	// The repository with image IDs to be listed.
	RepositoryName *string
	// The AWS account ID associated with the registry that contains the repository in
	// which to list images. If you do not specify a registry, the default registry is
	// assumed.
	RegistryId *string
}

type ListImagesOutput struct {
	// The nextToken value to include in a future ListImages request. When the results
	// of a ListImages request exceed maxResults, this value can be used to retrieve
	// the next page of results. This value is null when there are no more results to
	// return.
	NextToken *string
	// The list of image IDs for the requested repository.
	ImageIds []*types.ImageIdentifier

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpListImagesMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpListImages{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpListImages{}, middleware.After)
}

func newServiceMetadataMiddleware_opListImages(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ecr",
		OperationName: "ListImages",
	}
}