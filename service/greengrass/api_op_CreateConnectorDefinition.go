// Code generated by smithy-go-codegen DO NOT EDIT.

package greengrass

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/greengrass/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a connector definition. You may provide the initial version of the
// connector definition now or use ''CreateConnectorDefinitionVersion'' at a later
// time.
func (c *Client) CreateConnectorDefinition(ctx context.Context, params *CreateConnectorDefinitionInput, optFns ...func(*Options)) (*CreateConnectorDefinitionOutput, error) {
	stack := middleware.NewStack("CreateConnectorDefinition", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsRestjson1_serdeOpCreateConnectorDefinitionMiddlewares(stack)
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
	addOpCreateConnectorDefinitionValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateConnectorDefinition(options.Region), middleware.Before)

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
			OperationName: "CreateConnectorDefinition",
			Err:           err,
		}
	}
	out := result.(*CreateConnectorDefinitionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateConnectorDefinitionInput struct {
	// The name of the connector definition.
	Name *string
	// Tag(s) to add to the new resource.
	Tags map[string]*string
	// Information about the initial version of the connector definition.
	InitialVersion *types.ConnectorDefinitionVersion
	// A client token used to correlate requests and responses.
	AmznClientToken *string
}

type CreateConnectorDefinitionOutput struct {
	// The ID of the definition.
	Id *string
	// The time, in milliseconds since the epoch, when the definition was last updated.
	LastUpdatedTimestamp *string
	// The name of the definition.
	Name *string
	// The ARN of the definition.
	Arn *string
	// The time, in milliseconds since the epoch, when the definition was created.
	CreationTimestamp *string
	// The ID of the latest version associated with the definition.
	LatestVersion *string
	// The ARN of the latest version associated with the definition.
	LatestVersionArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsRestjson1_serdeOpCreateConnectorDefinitionMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsRestjson1_serializeOpCreateConnectorDefinition{}, middleware.After)
	stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateConnectorDefinition{}, middleware.After)
}

func newServiceMetadataMiddleware_opCreateConnectorDefinition(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "greengrass",
		OperationName: "CreateConnectorDefinition",
	}
}