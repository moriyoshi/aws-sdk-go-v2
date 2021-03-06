// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Deletes a reference (route) to a prefix list in a specified transit gateway
// route table.
func (c *Client) DeleteTransitGatewayPrefixListReference(ctx context.Context, params *DeleteTransitGatewayPrefixListReferenceInput, optFns ...func(*Options)) (*DeleteTransitGatewayPrefixListReferenceOutput, error) {
	if params == nil {
		params = &DeleteTransitGatewayPrefixListReferenceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteTransitGatewayPrefixListReference", params, optFns, addOperationDeleteTransitGatewayPrefixListReferenceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteTransitGatewayPrefixListReferenceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteTransitGatewayPrefixListReferenceInput struct {

	// The ID of the prefix list.
	//
	// This member is required.
	PrefixListId *string

	// The ID of the route table.
	//
	// This member is required.
	TransitGatewayRouteTableId *string

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation. Otherwise, it is
	// UnauthorizedOperation.
	DryRun *bool
}

type DeleteTransitGatewayPrefixListReferenceOutput struct {

	// Information about the deleted prefix list reference.
	TransitGatewayPrefixListReference *types.TransitGatewayPrefixListReference

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteTransitGatewayPrefixListReferenceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpDeleteTransitGatewayPrefixListReference{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpDeleteTransitGatewayPrefixListReference{}, middleware.After)
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
	addOpDeleteTransitGatewayPrefixListReferenceValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteTransitGatewayPrefixListReference(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteTransitGatewayPrefixListReference(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "DeleteTransitGatewayPrefixListReference",
	}
}
