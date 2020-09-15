// Code generated by smithy-go-codegen DO NOT EDIT.

package worklink

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/worklink/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves a list of website authorization providers associated with a specified
// fleet.
func (c *Client) ListWebsiteAuthorizationProviders(ctx context.Context, params *ListWebsiteAuthorizationProvidersInput, optFns ...func(*Options)) (*ListWebsiteAuthorizationProvidersOutput, error) {
	stack := middleware.NewStack("ListWebsiteAuthorizationProviders", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpListWebsiteAuthorizationProvidersMiddlewares(stack)
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
	addOpListWebsiteAuthorizationProvidersValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opListWebsiteAuthorizationProviders(options.Region), middleware.Before)

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
			OperationName: "ListWebsiteAuthorizationProviders",
			Err:           err,
		}
	}
	out := result.(*ListWebsiteAuthorizationProvidersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListWebsiteAuthorizationProvidersInput struct {
	// The ARN of the fleet.
	FleetArn *string
	// The maximum number of results to be included in the next page.
	MaxResults *int32
	// The pagination token to use to retrieve the next page of results for this
	// operation. If this value is null, it retrieves the first page.
	NextToken *string
}

type ListWebsiteAuthorizationProvidersOutput struct {
	// The website authorization providers.
	WebsiteAuthorizationProviders []*types.WebsiteAuthorizationProviderSummary
	// The pagination token to use to retrieve the next page of results for this
	// operation. If this value is null, it retrieves the first page.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpListWebsiteAuthorizationProvidersMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpListWebsiteAuthorizationProviders{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpListWebsiteAuthorizationProviders{}, middleware.After)
}

func newServiceMetadataMiddleware_opListWebsiteAuthorizationProviders(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "worklink",
		OperationName: "ListWebsiteAuthorizationProviders",
	}
}