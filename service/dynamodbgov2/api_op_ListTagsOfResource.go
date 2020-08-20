// Code generated by smithy-go-codegen DO NOT EDIT.

package dynamodbgov2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/dynamodbgov2/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// List all tags on an Amazon DynamoDB resource. You can call ListTagsOfResource up
// to 10 times per second, per account. For an overview on tagging DynamoDB
// resources, see Tagging for DynamoDB
// (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Tagging.html)
// in the Amazon DynamoDB Developer Guide.
func (c *Client) ListTagsOfResource(ctx context.Context, params *ListTagsOfResourceInput, optFns ...func(*Options)) (*ListTagsOfResourceOutput, error) {
	stack := middleware.NewStack("ListTagsOfResource", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson10_serdeOpListTagsOfResourceMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	v4.AddHTTPSignerMiddleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpListTagsOfResourceValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListTagsOfResource(options.Region), middleware.Before)

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, err
		}
	}
	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err := handler.Handle(ctx, params)
	if err != nil {
		return nil, &smithy.OperationError{
			ServiceID:     c.ServiceID(),
			OperationName: "ListTagsOfResource",
			Err:           err,
		}
	}
	out := result.(*ListTagsOfResourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListTagsOfResourceInput struct {
	// An optional string that, if supplied, must be copied from the output of a
	// previous call to ListTagOfResource. When provided in this manner, this API
	// fetches the next page of results.
	NextToken *string
	// The Amazon DynamoDB resource with tags to be listed. This value is an Amazon
	// Resource Name (ARN).
	ResourceArn *string
}

type ListTagsOfResourceOutput struct {
	// If this value is returned, there are additional results to be displayed. To
	// retrieve them, call ListTagsOfResource again, with NextToken set to this value.
	NextToken *string
	// The tags currently associated with the Amazon DynamoDB resource.
	Tags []*types.Tag

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson10_serdeOpListTagsOfResourceMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson10_serializeOpListTagsOfResource{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson10_deserializeOpListTagsOfResource{}, middleware.After)
}

func newServiceMetadataMiddleware_opListTagsOfResource(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:         region,
		ServiceName:    "DynamoDB GoV2",
		ServiceID:      "dynamodbgov2",
		EndpointPrefix: "dynamodbgov2",
		SigningName:    "dynamodb",
		OperationName:  "ListTagsOfResource",
	}
}
