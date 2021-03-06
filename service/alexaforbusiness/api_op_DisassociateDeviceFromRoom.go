// Code generated by smithy-go-codegen DO NOT EDIT.

package alexaforbusiness

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Disassociates a device from its current room. The device continues to be
// connected to the Wi-Fi network and is still registered to the account. The
// device settings and skills are removed from the room.
func (c *Client) DisassociateDeviceFromRoom(ctx context.Context, params *DisassociateDeviceFromRoomInput, optFns ...func(*Options)) (*DisassociateDeviceFromRoomOutput, error) {
	if params == nil {
		params = &DisassociateDeviceFromRoomInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DisassociateDeviceFromRoom", params, optFns, addOperationDisassociateDeviceFromRoomMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DisassociateDeviceFromRoomOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisassociateDeviceFromRoomInput struct {

	// The ARN of the device to disassociate from a room. Required.
	DeviceArn *string
}

type DisassociateDeviceFromRoomOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDisassociateDeviceFromRoomMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDisassociateDeviceFromRoom{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDisassociateDeviceFromRoom{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opDisassociateDeviceFromRoom(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDisassociateDeviceFromRoom(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "a4b",
		OperationName: "DisassociateDeviceFromRoom",
	}
}
