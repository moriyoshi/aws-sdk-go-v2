// Code generated by smithy-go-codegen DO NOT EDIT.

package iot

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iot/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Describes a registered CA certificate.
func (c *Client) DescribeCACertificate(ctx context.Context, params *DescribeCACertificateInput, optFns ...func(*Options)) (*DescribeCACertificateOutput, error) {
	if params == nil {
		params = &DescribeCACertificateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeCACertificate", params, optFns, addOperationDescribeCACertificateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeCACertificateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input for the DescribeCACertificate operation.
type DescribeCACertificateInput struct {

	// The CA certificate identifier.
	//
	// This member is required.
	CertificateId *string
}

// The output from the DescribeCACertificate operation.
type DescribeCACertificateOutput struct {

	// The CA certificate description.
	CertificateDescription *types.CACertificateDescription

	// Information about the registration configuration.
	RegistrationConfig *types.RegistrationConfig

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeCACertificateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeCACertificate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeCACertificate{}, middleware.After)
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
	addOpDescribeCACertificateValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeCACertificate(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeCACertificate(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "execute-api",
		OperationName: "DescribeCACertificate",
	}
}
