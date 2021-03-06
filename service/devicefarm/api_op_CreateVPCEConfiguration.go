// Code generated by smithy-go-codegen DO NOT EDIT.

package devicefarm

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/devicefarm/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Creates a configuration record in Device Farm for your Amazon Virtual Private
// Cloud (VPC) endpoint.
func (c *Client) CreateVPCEConfiguration(ctx context.Context, params *CreateVPCEConfigurationInput, optFns ...func(*Options)) (*CreateVPCEConfigurationOutput, error) {
	if params == nil {
		params = &CreateVPCEConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateVPCEConfiguration", params, optFns, addOperationCreateVPCEConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateVPCEConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateVPCEConfigurationInput struct {

	// The DNS name of the service running in your VPC that you want Device Farm to
	// test.
	//
	// This member is required.
	ServiceDnsName *string

	// The friendly name you give to your VPC endpoint configuration, to manage your
	// configurations more easily.
	//
	// This member is required.
	VpceConfigurationName *string

	// The name of the VPC endpoint service running in your AWS account that you want
	// Device Farm to test.
	//
	// This member is required.
	VpceServiceName *string

	// An optional description that provides details about your VPC endpoint
	// configuration.
	VpceConfigurationDescription *string
}

type CreateVPCEConfigurationOutput struct {

	// An object that contains information about your VPC endpoint configuration.
	VpceConfiguration *types.VPCEConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationCreateVPCEConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateVPCEConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateVPCEConfiguration{}, middleware.After)
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
	addOpCreateVPCEConfigurationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opCreateVPCEConfiguration(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opCreateVPCEConfiguration(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "devicefarm",
		OperationName: "CreateVPCEConfiguration",
	}
}
