// Code generated by smithy-go-codegen DO NOT EDIT.

package kms

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Enables automatic rotation of the key material
// (https://docs.aws.amazon.com/kms/latest/developerguide/rotate-keys.html) for the
// specified symmetric customer master key (CMK). You cannot perform this operation
// on a CMK in a different AWS account. You cannot enable automatic rotation of
// asymmetric CMKs, CMKs with imported key material, or CMKs in a custom key store
// (https://docs.aws.amazon.com/kms/latest/developerguide/custom-key-store-overview.html).
// The CMK that you use for this operation must be in a compatible key state. For
// details, see How Key State Affects Use of a Customer Master Key
// (https://docs.aws.amazon.com/kms/latest/developerguide/key-state.html) in the
// AWS Key Management Service Developer Guide.
func (c *Client) EnableKeyRotation(ctx context.Context, params *EnableKeyRotationInput, optFns ...func(*Options)) (*EnableKeyRotationOutput, error) {
	stack := middleware.NewStack("EnableKeyRotation", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpEnableKeyRotationMiddlewares(stack)
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
	addOpEnableKeyRotationValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opEnableKeyRotation(options.Region), middleware.Before)

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
			OperationName: "EnableKeyRotation",
			Err:           err,
		}
	}
	out := result.(*EnableKeyRotationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type EnableKeyRotationInput struct {
	// Identifies a symmetric customer master key (CMK). You cannot enable automatic
	// rotation of asymmetric CMKs, CMKs with imported key material, or CMKs in a
	// custom key store
	// (https://docs.aws.amazon.com/kms/latest/developerguide/custom-key-store-overview.html).
	// <p>Specify the key ID or the Amazon Resource Name (ARN) of the CMK.</p> <p>For
	// example:</p> <ul> <li> <p>Key ID:
	// <code>1234abcd-12ab-34cd-56ef-1234567890ab</code> </p> </li> <li> <p>Key ARN:
	// <code>arn:aws:kms:us-east-2:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab</code>
	// </p> </li> </ul> <p>To get the key ID and key ARN for a CMK, use <a>ListKeys</a>
	// or <a>DescribeKey</a>.</p>
	KeyId *string
}

type EnableKeyRotationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpEnableKeyRotationMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpEnableKeyRotation{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpEnableKeyRotation{}, middleware.After)
}

func newServiceMetadataMiddleware_opEnableKeyRotation(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kms",
		OperationName: "EnableKeyRotation",
	}
}