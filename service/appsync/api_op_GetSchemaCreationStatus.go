// Code generated by smithy-go-codegen DO NOT EDIT.

package appsync

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/appsync/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Retrieves the current status of a schema creation operation.
func (c *Client) GetSchemaCreationStatus(ctx context.Context, params *GetSchemaCreationStatusInput, optFns ...func(*Options)) (*GetSchemaCreationStatusOutput, error) {
	stack := middleware.NewStack("GetSchemaCreationStatus", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetSchemaCreationStatusMiddlewares(stack)
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
	addOpGetSchemaCreationStatusValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetSchemaCreationStatus(options.Region), middleware.Before)

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
			OperationName: "GetSchemaCreationStatus",
			Err:           err,
		}
	}
	out := result.(*GetSchemaCreationStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetSchemaCreationStatusInput struct {
	// The API ID.
	ApiId *string
}

type GetSchemaCreationStatusOutput struct {
	// Detailed information about the status of the schema creation operation.
	Details *string
	// The current state of the schema (PROCESSING, FAILED, SUCCESS, or
	// NOT_APPLICABLE). When the schema is in the ACTIVE state, you can add data.
	Status types.SchemaStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetSchemaCreationStatusMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetSchemaCreationStatus{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetSchemaCreationStatus{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetSchemaCreationStatus(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "appsync",
		OperationName: "GetSchemaCreationStatus",
	}
}