// Code generated by smithy-go-codegen DO NOT EDIT.

package frauddetector

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/frauddetector/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets all entity types or a specific entity type if a name is specified. This is
// a paginated API. If you provide a null maxResults, this action retrieves a
// maximum of 10 records per page. If you provide a maxResults, the value must be
// between 5 and 10. To get the next page results, provide the pagination token
// from the GetEntityTypesResponse as part of your request. A null pagination token
// fetches the records from the beginning.
func (c *Client) GetEntityTypes(ctx context.Context, params *GetEntityTypesInput, optFns ...func(*Options)) (*GetEntityTypesOutput, error) {
	if params == nil {
		params = &GetEntityTypesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetEntityTypes", params, optFns, addOperationGetEntityTypesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetEntityTypesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetEntityTypesInput struct {

	// The maximum number of objects to return for the request.
	MaxResults *int32

	// The name.
	Name *string

	// The next token for the subsequent request.
	NextToken *string
}

type GetEntityTypesOutput struct {

	// An array of entity types.
	EntityTypes []*types.EntityType

	// The next page token.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetEntityTypesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetEntityTypes{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetEntityTypes{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetEntityTypes(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetEntityTypes(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "frauddetector",
		OperationName: "GetEntityTypes",
	}
}
