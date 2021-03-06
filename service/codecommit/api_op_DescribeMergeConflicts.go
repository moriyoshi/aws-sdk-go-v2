// Code generated by smithy-go-codegen DO NOT EDIT.

package codecommit

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codecommit/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns information about one or more merge conflicts in the attempted merge of
// two commit specifiers using the squash or three-way merge strategy. If the merge
// option for the attempted merge is specified as FAST_FORWARD_MERGE, an exception
// is thrown.
func (c *Client) DescribeMergeConflicts(ctx context.Context, params *DescribeMergeConflictsInput, optFns ...func(*Options)) (*DescribeMergeConflictsOutput, error) {
	if params == nil {
		params = &DescribeMergeConflictsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeMergeConflicts", params, optFns, addOperationDescribeMergeConflictsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeMergeConflictsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeMergeConflictsInput struct {

	// The branch, tag, HEAD, or other fully qualified reference used to identify a
	// commit (for example, a branch name or a full commit ID).
	//
	// This member is required.
	DestinationCommitSpecifier *string

	// The path of the target files used to describe the conflicts.
	//
	// This member is required.
	FilePath *string

	// The merge option or strategy you want to use to merge the code.
	//
	// This member is required.
	MergeOption types.MergeOptionTypeEnum

	// The name of the repository where you want to get information about a merge
	// conflict.
	//
	// This member is required.
	RepositoryName *string

	// The branch, tag, HEAD, or other fully qualified reference used to identify a
	// commit (for example, a branch name or a full commit ID).
	//
	// This member is required.
	SourceCommitSpecifier *string

	// The level of conflict detail to use. If unspecified, the default FILE_LEVEL is
	// used, which returns a not-mergeable result if the same file has differences in
	// both branches. If LINE_LEVEL is specified, a conflict is considered not
	// mergeable if the same file in both branches has differences on the same line.
	ConflictDetailLevel types.ConflictDetailLevelTypeEnum

	// Specifies which branch to use when resolving conflicts, or whether to attempt
	// automatically merging two versions of a file. The default is NONE, which
	// requires any conflicts to be resolved manually before the merge operation is
	// successful.
	ConflictResolutionStrategy types.ConflictResolutionStrategyTypeEnum

	// The maximum number of merge hunks to include in the output.
	MaxMergeHunks *int32

	// An enumeration token that, when provided in a request, returns the next batch of
	// the results.
	NextToken *string
}

type DescribeMergeConflictsOutput struct {

	// Contains metadata about the conflicts found in the merge.
	//
	// This member is required.
	ConflictMetadata *types.ConflictMetadata

	// The commit ID of the destination commit specifier that was used in the merge
	// evaluation.
	//
	// This member is required.
	DestinationCommitId *string

	// A list of merge hunks of the differences between the files or lines.
	//
	// This member is required.
	MergeHunks []*types.MergeHunk

	// The commit ID of the source commit specifier that was used in the merge
	// evaluation.
	//
	// This member is required.
	SourceCommitId *string

	// The commit ID of the merge base.
	BaseCommitId *string

	// An enumeration token that can be used in a request to return the next batch of
	// the results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationDescribeMergeConflictsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeMergeConflicts{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeMergeConflicts{}, middleware.After)
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
	addOpDescribeMergeConflictsValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeMergeConflicts(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opDescribeMergeConflicts(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codecommit",
		OperationName: "DescribeMergeConflicts",
	}
}
