// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/retry"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/gamelift/types"
	smithy "github.com/awslabs/smithy-go"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Updates settings for a game session queue, which determines how new game session
// requests in the queue are processed. To update settings, specify the queue name
// to be updated and provide the new settings. When updating destinations, provide
// a complete list of destinations. Learn more  Using Multi-Region Queues
// (https://docs.aws.amazon.com/gamelift/latest/developerguide/queues-intro.html)
// Related operations
//
//     * CreateGameSessionQueue ()
//
//     *
// DescribeGameSessionQueues ()
//
//     * UpdateGameSessionQueue ()
//
//     *
// DeleteGameSessionQueue ()
func (c *Client) UpdateGameSessionQueue(ctx context.Context, params *UpdateGameSessionQueueInput, optFns ...func(*Options)) (*UpdateGameSessionQueueOutput, error) {
	stack := middleware.NewStack("UpdateGameSessionQueue", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpUpdateGameSessionQueueMiddlewares(stack)
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
	addOpUpdateGameSessionQueueValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateGameSessionQueue(options.Region), middleware.Before)

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
			OperationName: "UpdateGameSessionQueue",
			Err:           err,
		}
	}
	out := result.(*UpdateGameSessionQueueOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input for a request action.
type UpdateGameSessionQueueInput struct {
	// A collection of latency policies to apply when processing game sessions
	// placement requests with player latency information. Multiple policies are
	// evaluated in order of the maximum latency value, starting with the lowest
	// latency values. With just one policy, the policy is enforced at the start of the
	// game session placement for the duration period. With multiple policies, each
	// policy is enforced consecutively for its duration period. For example, a queue
	// might enforce a 60-second policy followed by a 120-second policy, and then no
	// policy for the remainder of the placement. When updating policies, provide a
	// complete collection of policies.
	PlayerLatencyPolicies []*types.PlayerLatencyPolicy
	// The maximum time, in seconds, that a new game session placement request remains
	// in the queue. When a request exceeds this time, the game session placement
	// changes to a TIMED_OUT status.
	TimeoutInSeconds *int32
	// A list of fleets that can be used to fulfill game session placement requests in
	// the queue. Fleets are identified by either a fleet ARN or a fleet alias ARN.
	// Destinations are listed in default preference order. When updating this list,
	// provide a complete list of destinations.
	Destinations []*types.GameSessionQueueDestination
	// A descriptive label that is associated with game session queue. Queue names must
	// be unique within each Region. You can use either the queue ID or ARN value.
	Name *string
}

// Represents the returned data in response to a request action.
type UpdateGameSessionQueueOutput struct {
	// An object that describes the newly updated game session queue.
	GameSessionQueue *types.GameSessionQueue

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpUpdateGameSessionQueueMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateGameSessionQueue{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateGameSessionQueue{}, middleware.After)
}

func newServiceMetadataMiddleware_opUpdateGameSessionQueue(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "UpdateGameSessionQueue",
	}
}