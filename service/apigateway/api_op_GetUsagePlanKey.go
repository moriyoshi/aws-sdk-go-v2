// Code generated by smithy-go-codegen DO NOT EDIT.

package apigateway

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Gets a usage plan key of a given key identifier.
func (c *Client) GetUsagePlanKey(ctx context.Context, params *GetUsagePlanKeyInput, optFns ...func(*Options)) (*GetUsagePlanKeyOutput, error) {
	stack := middleware.NewStack("GetUsagePlanKey", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpGetUsagePlanKeyMiddlewares(stack)
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
	addOpGetUsagePlanKeyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetUsagePlanKey(options.Region), middleware.Before)
	addAcceptHeader(stack)

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
			OperationName: "GetUsagePlanKey",
			Err:           err,
		}
	}
	out := result.(*GetUsagePlanKeyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The GET request to get a usage plan key of a given key identifier.
type GetUsagePlanKeyInput struct {
	Template *bool
	Name     *string
	// [Required] The key Id of the to-be-retrieved UsagePlanKey () resource
	// representing a plan customer.
	KeyId            *string
	TemplateSkipList []*string
	Title            *string
	// [Required] The Id of the UsagePlan () resource representing the usage plan
	// containing the to-be-retrieved UsagePlanKey () resource representing a plan
	// customer.
	UsagePlanId *string
}

// Represents a usage plan key to identify a plan customer. To associate an API
// stage with a selected API key in a usage plan, you must create a UsagePlanKey
// resource to represent the selected ApiKey (). " Create and Use Usage Plans
// (https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-api-usage-plans.html)
type GetUsagePlanKeyOutput struct {
	// The name of a usage plan key.
	Name *string
	// The value of a usage plan key.
	Value *string
	// The Id of a usage plan key.
	Id *string
	// The type of a usage plan key. Currently, the valid key type is API_KEY.
	Type *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpGetUsagePlanKeyMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpGetUsagePlanKey{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpGetUsagePlanKey{}, middleware.After)
}

func newServiceMetadataMiddleware_opGetUsagePlanKey(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "apigateway",
		OperationName: "GetUsagePlanKey",
	}
}
