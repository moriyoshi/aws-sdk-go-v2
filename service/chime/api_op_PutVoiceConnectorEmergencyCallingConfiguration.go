// Code generated by smithy-go-codegen DO NOT EDIT.

package chime

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/chime/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Puts emergency calling configuration details to the specified Amazon Chime Voice
// Connector, such as emergency phone numbers and calling countries. Origination
// and termination settings must be enabled for the Amazon Chime Voice Connector
// before emergency calling can be configured.
func (c *Client) PutVoiceConnectorEmergencyCallingConfiguration(ctx context.Context, params *PutVoiceConnectorEmergencyCallingConfigurationInput, optFns ...func(*Options)) (*PutVoiceConnectorEmergencyCallingConfigurationOutput, error) {
	if params == nil {
		params = &PutVoiceConnectorEmergencyCallingConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutVoiceConnectorEmergencyCallingConfiguration", params, optFns, addOperationPutVoiceConnectorEmergencyCallingConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutVoiceConnectorEmergencyCallingConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutVoiceConnectorEmergencyCallingConfigurationInput struct {

	// The emergency calling configuration details.
	//
	// This member is required.
	EmergencyCallingConfiguration *types.EmergencyCallingConfiguration

	// The Amazon Chime Voice Connector ID.
	//
	// This member is required.
	VoiceConnectorId *string
}

type PutVoiceConnectorEmergencyCallingConfigurationOutput struct {

	// The emergency calling configuration details.
	EmergencyCallingConfiguration *types.EmergencyCallingConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationPutVoiceConnectorEmergencyCallingConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpPutVoiceConnectorEmergencyCallingConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpPutVoiceConnectorEmergencyCallingConfiguration{}, middleware.After)
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
	addOpPutVoiceConnectorEmergencyCallingConfigurationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPutVoiceConnectorEmergencyCallingConfiguration(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opPutVoiceConnectorEmergencyCallingConfiguration(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "PutVoiceConnectorEmergencyCallingConfiguration",
	}
}
