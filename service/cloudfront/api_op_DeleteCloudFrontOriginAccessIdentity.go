// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudfront

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Delete an origin access identity.
func (c *Client) DeleteCloudFrontOriginAccessIdentity(ctx context.Context, params *DeleteCloudFrontOriginAccessIdentityInput, optFns ...func(*Options)) (*DeleteCloudFrontOriginAccessIdentityOutput, error) {
	if params == nil {
		params = &DeleteCloudFrontOriginAccessIdentityInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteCloudFrontOriginAccessIdentity", params, optFns, addOperationDeleteCloudFrontOriginAccessIdentityMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteCloudFrontOriginAccessIdentityOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Deletes a origin access identity.
type DeleteCloudFrontOriginAccessIdentityInput struct {

	// The origin access identity's ID.
	//
	// This member is required.
	Id *string

	// The value of the ETag header you received from a previous GET or PUT request.
	// For example: E2QWRUHAPOMQZL.
	IfMatch *string
}

type DeleteCloudFrontOriginAccessIdentityOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteCloudFrontOriginAccessIdentityMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpDeleteCloudFrontOriginAccessIdentity{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpDeleteCloudFrontOriginAccessIdentity{}, middleware.After)
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
	addOpDeleteCloudFrontOriginAccessIdentityValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteCloudFrontOriginAccessIdentity(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteCloudFrontOriginAccessIdentity(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudfront",
		OperationName: "DeleteCloudFrontOriginAccessIdentity",
	}
}
