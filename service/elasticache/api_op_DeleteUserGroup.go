// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticache

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticache/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// For Redis engine version 6.04 onwards: Deletes a ser group. The user group must
// first be disassociated from the replcation group before it can be deleted. For
// more information, see Using Role Based Access Control (RBAC)
// (http://docs.aws.amazon.com/AmazonElastiCache/latest/red-ug/Clusters.RBAC.html).
func (c *Client) DeleteUserGroup(ctx context.Context, params *DeleteUserGroupInput, optFns ...func(*Options)) (*DeleteUserGroupOutput, error) {
	if params == nil {
		params = &DeleteUserGroupInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteUserGroup", params, optFns, addOperationDeleteUserGroupMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteUserGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteUserGroupInput struct {

	// The ID of the user group.
	//
	// This member is required.
	UserGroupId *string
}

type DeleteUserGroupOutput struct {

	// The Amazon Resource Name (ARN) of the user group.
	ARN *string

	// Must be Redis.
	Engine *string

	// A list of updates being applied to the user groups.
	PendingChanges *types.UserGroupPendingChanges

	// A list of replication groups that the user group can access.
	ReplicationGroups []*string

	// Indicates user group status. Can be "creating", "active", "modifying",
	// "deleting".
	Status *string

	// The ID of the user group.
	UserGroupId *string

	// The list of user IDs that belong to the user group.
	UserIds []*string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDeleteUserGroupMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDeleteUserGroup{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDeleteUserGroup{}, middleware.After)
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
	addOpDeleteUserGroupValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteUserGroup(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDeleteUserGroup(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticache",
		OperationName: "DeleteUserGroup",
	}
}
