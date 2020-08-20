// Code generated by smithy-go-codegen DO NOT EDIT.

package dynamodbgov2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/dynamodbgov2/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// UpdateContinuousBackups enables or disables point in time recovery for the
// specified table. A successful UpdateContinuousBackups call returns the current
// ContinuousBackupsDescription. Continuous backups are ENABLED on all tables at
// table creation. If point in time recovery is enabled, PointInTimeRecoveryStatus
// will be set to ENABLED. Once continuous backups and point in time recovery are
// enabled, you can restore to any point in time within EarliestRestorableDateTime
// and LatestRestorableDateTime. LatestRestorableDateTime is typically 5 minutes
// before the current time. You can restore your table to any point in time during
// the last 35 days.
func (c *Client) UpdateContinuousBackups(ctx context.Context, params *UpdateContinuousBackupsInput, optFns ...func(*Options)) (*UpdateContinuousBackupsOutput, error) {
	stack := middleware.NewStack("UpdateContinuousBackups", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson10_serdeOpUpdateContinuousBackupsMiddlewares(stack)
	awsmiddleware.AddRequestInvocationIDMiddleware(stack)
	smithyhttp.AddContentLengthMiddleware(stack)
	AddResolveEndpointMiddleware(stack, options)
	v4.AddComputePayloadSHA256Middleware(stack)
	retry.AddRetryMiddlewares(stack, options)
	v4.AddHTTPSignerMiddleware(stack, options)
	awsmiddleware.AddAttemptClockSkewMiddleware(stack)
	smithyhttp.AddErrorCloseResponseBodyMiddleware(stack)
	smithyhttp.AddCloseResponseBodyMiddleware(stack)
	addOpUpdateContinuousBackupsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateContinuousBackups(options.Region), middleware.Before)

	for _, fn := range options.APIOptions {
		if err := fn(stack); err != nil {
			return nil, err
		}
	}
	handler := middleware.DecorateHandler(smithyhttp.NewClientHandler(options.HTTPClient), stack)
	result, metadata, err := handler.Handle(ctx, params)
	if err != nil {
		return nil, &smithy.OperationError{
			ServiceID:     c.ServiceID(),
			OperationName: "UpdateContinuousBackups",
			Err:           err,
		}
	}
	out := result.(*UpdateContinuousBackupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateContinuousBackupsInput struct {
	// Represents the settings used to enable point in time recovery.
	PointInTimeRecoverySpecification *types.PointInTimeRecoverySpecification
	// The name of the table.
	TableName *string
}

type UpdateContinuousBackupsOutput struct {
	// Represents the continuous backups and point in time recovery settings on the
	// table.
	ContinuousBackupsDescription *types.ContinuousBackupsDescription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson10_serdeOpUpdateContinuousBackupsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson10_serializeOpUpdateContinuousBackups{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson10_deserializeOpUpdateContinuousBackups{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateContinuousBackups(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:         region,
		ServiceName:    "DynamoDB GoV2",
		ServiceID:      "dynamodbgov2",
		EndpointPrefix: "dynamodbgov2",
		SigningName:    "dynamodb",
		OperationName:  "UpdateContinuousBackups",
	}
}
