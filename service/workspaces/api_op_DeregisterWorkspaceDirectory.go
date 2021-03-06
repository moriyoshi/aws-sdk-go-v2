// Code generated by smithy-go-codegen DO NOT EDIT.

package workspaces

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deregisters the specified directory. This operation is asynchronous and returns
// before the WorkSpace directory is deregistered. If any WorkSpaces are registered
// to this directory, you must remove them before you can deregister the directory.
func (c *Client) DeregisterWorkspaceDirectory(ctx context.Context, params *DeregisterWorkspaceDirectoryInput, optFns ...func(*Options)) (*DeregisterWorkspaceDirectoryOutput, error) {
	if params == nil {
		params = &DeregisterWorkspaceDirectoryInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeregisterWorkspaceDirectory", params, optFns, addOperationDeregisterWorkspaceDirectoryMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeregisterWorkspaceDirectoryOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeregisterWorkspaceDirectoryInput struct {

	// The identifier of the directory. If any WorkSpaces are registered to this
	// directory, you must remove them before you deregister the directory, or you will
	// receive an OperationNotSupportedException error.
	//
	// This member is required.
	DirectoryId *string
}

type DeregisterWorkspaceDirectoryOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeregisterWorkspaceDirectoryMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDeregisterWorkspaceDirectory{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDeregisterWorkspaceDirectory{}, middleware.After)
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
	addOpDeregisterWorkspaceDirectoryValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeregisterWorkspaceDirectory(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeregisterWorkspaceDirectory(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workspaces",
		OperationName: "DeregisterWorkspaceDirectory",
	}
}
