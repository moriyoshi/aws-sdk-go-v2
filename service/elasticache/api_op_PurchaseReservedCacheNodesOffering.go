// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticache

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticache/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Allows you to purchase a reserved cache node offering.
func (c *Client) PurchaseReservedCacheNodesOffering(ctx context.Context, params *PurchaseReservedCacheNodesOfferingInput, optFns ...func(*Options)) (*PurchaseReservedCacheNodesOfferingOutput, error) {
	stack := middleware.NewStack("PurchaseReservedCacheNodesOffering", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsquery_serdeOpPurchaseReservedCacheNodesOfferingMiddlewares(stack)
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
	addOpPurchaseReservedCacheNodesOfferingValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opPurchaseReservedCacheNodesOffering(options.Region), middleware.Before)

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
			OperationName: "PurchaseReservedCacheNodesOffering",
			Err:           err,
		}
	}
	out := result.(*PurchaseReservedCacheNodesOfferingOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a PurchaseReservedCacheNodesOffering operation.
type PurchaseReservedCacheNodesOfferingInput struct {
	// The ID of the reserved cache node offering to purchase. Example:
	// 438012d3-4052-4cc7-b2e3-8d3372e0e706
	ReservedCacheNodesOfferingId *string
	// The number of cache node instances to reserve. Default: 1
	CacheNodeCount *int32
	// A customer-specified identifier to track this reservation. The Reserved Cache
	// Node ID is an unique customer-specified identifier to track this reservation. If
	// this parameter is not specified, ElastiCache automatically generates an
	// identifier for the reservation. Example: myreservationID
	ReservedCacheNodeId *string
}

type PurchaseReservedCacheNodesOfferingOutput struct {
	// Represents the output of a PurchaseReservedCacheNodesOffering operation.
	ReservedCacheNode *types.ReservedCacheNode

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsquery_serdeOpPurchaseReservedCacheNodesOfferingMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsquery_serializeOpPurchaseReservedCacheNodesOffering{}, middleware.After)
	stack.Deserialize.Add(&awsAwsquery_deserializeOpPurchaseReservedCacheNodesOffering{}, middleware.After)
}

func newServiceMetadataMiddleware_opPurchaseReservedCacheNodesOffering(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticache",
		OperationName: "PurchaseReservedCacheNodesOffering",
	}
}