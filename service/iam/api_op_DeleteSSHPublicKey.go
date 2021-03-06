// Code generated by smithy-go-codegen DO NOT EDIT.

package iam

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes the specified SSH public key. The SSH public key deleted by this
// operation is used only for authenticating the associated IAM user to an AWS
// CodeCommit repository. For more information about using SSH keys to authenticate
// to an AWS CodeCommit repository, see Set up AWS CodeCommit for SSH Connections
// (https://docs.aws.amazon.com/codecommit/latest/userguide/setting-up-credentials-ssh.html)
// in the AWS CodeCommit User Guide.
func (c *Client) DeleteSSHPublicKey(ctx context.Context, params *DeleteSSHPublicKeyInput, optFns ...func(*Options)) (*DeleteSSHPublicKeyOutput, error) {
	if params == nil {
		params = &DeleteSSHPublicKeyInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteSSHPublicKey", params, optFns, addOperationDeleteSSHPublicKeyMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteSSHPublicKeyOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteSSHPublicKeyInput struct {

	// The unique identifier for the SSH public key. This parameter allows (through its
	// regex pattern (http://wikipedia.org/wiki/regex)) a string of characters that can
	// consist of any upper or lowercased letter or digit.
	//
	// This member is required.
	SSHPublicKeyId *string

	// The name of the IAM user associated with the SSH public key. This parameter
	// allows (through its regex pattern (http://wikipedia.org/wiki/regex)) a string of
	// characters consisting of upper and lowercase alphanumeric characters with no
	// spaces. You can also include any of the following characters: _+=,.@-
	//
	// This member is required.
	UserName *string
}

type DeleteSSHPublicKeyOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteSSHPublicKeyMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDeleteSSHPublicKey{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDeleteSSHPublicKey{}, middleware.After)
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
	addOpDeleteSSHPublicKeyValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteSSHPublicKey(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteSSHPublicKey(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iam",
		OperationName: "DeleteSSHPublicKey",
	}
}
