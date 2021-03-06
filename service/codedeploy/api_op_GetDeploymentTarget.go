// Code generated by smithy-go-codegen DO NOT EDIT.

package codedeploy

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codedeploy/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns information about a deployment target.
func (c *Client) GetDeploymentTarget(ctx context.Context, params *GetDeploymentTargetInput, optFns ...func(*Options)) (*GetDeploymentTargetOutput, error) {
	if params == nil {
		params = &GetDeploymentTargetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetDeploymentTarget", params, optFns, addOperationGetDeploymentTargetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetDeploymentTargetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetDeploymentTargetInput struct {

	// The unique ID of a deployment.
	DeploymentId *string

	// The unique ID of a deployment target.
	TargetId *string
}

type GetDeploymentTargetOutput struct {

	// A deployment target that contains information about a deployment such as its
	// status, lifecycle events, and when it was last updated. It also contains
	// metadata about the deployment target. The deployment target metadata depends on
	// the deployment target's type (instanceTarget, lambdaTarget, or ecsTarget).
	DeploymentTarget *types.DeploymentTarget

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetDeploymentTargetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetDeploymentTarget{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetDeploymentTarget{}, middleware.After)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opGetDeploymentTarget(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opGetDeploymentTarget(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codedeploy",
		OperationName: "GetDeploymentTarget",
	}
}
