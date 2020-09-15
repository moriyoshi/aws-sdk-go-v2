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

// Retrieves build resources for all builds associated with the AWS account in use.
// You can limit results to builds that are in a specific status by using the
// Status parameter. Use the pagination parameters to retrieve results in a set of
// sequential pages. Build resources are not listed in any particular order. Learn
// more  Upload a Custom Server Build
// (https://docs.aws.amazon.com/gamelift/latest/developerguide/gamelift-build-intro.html)
// Related operations
//
//     * CreateBuild ()
//
//     * ListBuilds ()
//
//     *
// DescribeBuild ()
//
//     * UpdateBuild ()
//
//     * DeleteBuild ()
func (c *Client) ListBuilds(ctx context.Context, params *ListBuildsInput, optFns ...func(*Options)) (*ListBuildsOutput, error) {
	stack := middleware.NewStack("ListBuilds", smithyhttp.NewStackRequest)
	options := c.options.Copy()
	for _, fn := range optFns {
		fn(&options)
	}
	addawsAwsjson11_serdeOpListBuildsMiddlewares(stack)
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
	stack.Initialize.Add(newServiceMetadataMiddleware_opListBuilds(options.Region), middleware.Before)

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
			OperationName: "ListBuilds",
			Err:           err,
		}
	}
	out := result.(*ListBuildsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input for a request action.
type ListBuildsInput struct {
	// The maximum number of results to return. Use this parameter with NextToken to
	// get results as a set of sequential pages.
	Limit *int32
	// Build status to filter results by. To retrieve all builds, leave this parameter
	// empty. Possible build statuses include the following:
	//
	//     * INITIALIZED -- A
	// new build has been defined, but no files have been uploaded. You cannot create
	// fleets for builds that are in this status. When a build is successfully created,
	// the build status is set to this value.
	//
	//     * READY -- The game build has been
	// successfully uploaded. You can now create new fleets for this build.
	//
	//     *
	// FAILED -- The game build upload failed. You cannot create new fleets for this
	// build.
	Status types.BuildStatus
	// Token that indicates the start of the next sequential page of results. Use the
	// token that is returned with a previous call to this action. To start at the
	// beginning of the result set, do not specify a value.
	NextToken *string
}

// Represents the returned data in response to a request action.
type ListBuildsOutput struct {
	// A collection of build resources that match the request.
	Builds []*types.Build
	// Token that indicates where to resume retrieving results on the next call to this
	// action. If no token is returned, these results represent the end of the list.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addawsAwsjson11_serdeOpListBuildsMiddlewares(stack *middleware.Stack) {
	stack.Serialize.Add(&awsAwsjson11_serializeOpListBuilds{}, middleware.After)
	stack.Deserialize.Add(&awsAwsjson11_deserializeOpListBuilds{}, middleware.After)
}

func newServiceMetadataMiddleware_opListBuilds(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "ListBuilds",
	}
}