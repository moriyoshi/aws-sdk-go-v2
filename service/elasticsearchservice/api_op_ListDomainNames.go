// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticsearchservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns the name of all Elasticsearch domains owned by the current user's
// account.
func (c *Client) ListDomainNames(ctx context.Context, params *ListDomainNamesInput, optFns ...func(*Options)) (*ListDomainNamesOutput, error) {
	if params == nil {
		params = &ListDomainNamesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListDomainNames", params, optFns, addOperationListDomainNamesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListDomainNamesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListDomainNamesInput struct {
}

// The result of a ListDomainNames operation. Contains the names of all
// Elasticsearch domains owned by this account.
type ListDomainNamesOutput struct {

	// List of Elasticsearch domain names.
	DomainNames []*types.DomainInfo

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationListDomainNamesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListDomainNames{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListDomainNames{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListDomainNames(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opListDomainNames(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "es",
		OperationName: "ListDomainNames",
	}
}
