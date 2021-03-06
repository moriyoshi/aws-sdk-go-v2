// Code generated by smithy-go-codegen DO NOT EDIT.

package workspaces

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Disassociates the specified IP access control group from the specified
// directory.
func (c *Client) DisassociateIpGroups(ctx context.Context, params *DisassociateIpGroupsInput, optFns ...func(*Options)) (*DisassociateIpGroupsOutput, error) {
	if params == nil {
		params = &DisassociateIpGroupsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DisassociateIpGroups", params, optFns, addOperationDisassociateIpGroupsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DisassociateIpGroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisassociateIpGroupsInput struct {

	// The identifier of the directory.
	//
	// This member is required.
	DirectoryId *string

	// The identifiers of one or more IP access control groups.
	//
	// This member is required.
	GroupIds []*string
}

type DisassociateIpGroupsOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDisassociateIpGroupsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDisassociateIpGroups{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDisassociateIpGroups{}, middleware.After)
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
	addOpDisassociateIpGroupsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDisassociateIpGroups(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDisassociateIpGroups(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "workspaces",
		OperationName: "DisassociateIpGroups",
	}
}
