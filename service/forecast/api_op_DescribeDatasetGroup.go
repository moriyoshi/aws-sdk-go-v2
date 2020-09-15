// Code generated by smithy-go-codegen DO NOT EDIT.

package forecast

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/forecast/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"time"
)

// Describes a dataset group created using the CreateDatasetGroup () operation. In
// addition to listing the parameters provided in the CreateDatasetGroup request,
// this operation includes the following properties:
//
//     * DatasetArns - The
// datasets belonging to the group.
//
//     * CreationTime
//
//     *
// LastModificationTime
//
//     * Status
func (c *Client) DescribeDatasetGroup(ctx context.Context, params *DescribeDatasetGroupInput, optFns ...func(*Options)) (*DescribeDatasetGroupOutput, error) {
	stack := middleware.NewStack("DescribeDatasetGroup", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpDescribeDatasetGroupMiddlewares(stack)
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
	addOpDescribeDatasetGroupValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeDatasetGroup(options.Region), middleware.Before)

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
			OperationName: "DescribeDatasetGroup",
			Err:           err,
		}
	}
	out := result.(*DescribeDatasetGroupOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeDatasetGroupInput struct {
	// The Amazon Resource Name (ARN) of the dataset group.
	DatasetGroupArn *string
}

type DescribeDatasetGroupOutput struct {
	// When the dataset group was created or last updated from a call to the
	// UpdateDatasetGroup () operation. While the dataset group is being updated,
	// LastModificationTime is the current time of the DescribeDatasetGroup call.
	LastModificationTime *time.Time
	// The status of the dataset group. States include:
	//
	//     * ACTIVE
	//
	//     *
	// CREATE_PENDING, CREATE_IN_PROGRESS, CREATE_FAILED
	//
	//     * DELETE_PENDING,
	// DELETE_IN_PROGRESS, DELETE_FAILED
	//
	//     * UPDATE_PENDING, UPDATE_IN_PROGRESS,
	// UPDATE_FAILED
	//
	// The UPDATE states apply when you call the UpdateDatasetGroup ()
	// operation. The Status of the dataset group must be ACTIVE before you can use the
	// dataset group to create a predictor.
	Status *string
	// An array of Amazon Resource Names (ARNs) of the datasets contained in the
	// dataset group.
	DatasetArns []*string
	// The name of the dataset group.
	DatasetGroupName *string
	// The ARN of the dataset group.
	DatasetGroupArn *string
	// The domain associated with the dataset group.
	Domain types.Domain
	// When the dataset group was created.
	CreationTime *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpDescribeDatasetGroupMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeDatasetGroup{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeDatasetGroup{}, middleware.After)
}

func newServiceMetadataMiddleware_opDescribeDatasetGroup(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "forecast",
		OperationName: "DescribeDatasetGroup",
	}
}