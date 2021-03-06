// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticsearchservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticsearchservice/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Modifies the cluster configuration of the specified Elasticsearch domain,
// setting as setting the instance type and the number of instances.
func (c *Client) UpdateElasticsearchDomainConfig(ctx context.Context, params *UpdateElasticsearchDomainConfigInput, optFns ...func(*Options)) (*UpdateElasticsearchDomainConfigOutput, error) {
	if params == nil {
		params = &UpdateElasticsearchDomainConfigInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateElasticsearchDomainConfig", params, optFns, addOperationUpdateElasticsearchDomainConfigMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateElasticsearchDomainConfigOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Container for the parameters to the UpdateElasticsearchDomain operation.
// Specifies the type and number of instances in the domain cluster.
type UpdateElasticsearchDomainConfigInput struct {

	// The name of the Elasticsearch domain that you are updating.
	//
	// This member is required.
	DomainName *string

	// IAM access policy as a JSON-formatted string.
	AccessPolicies *string

	// Modifies the advanced option to allow references to indices in an HTTP request
	// body. Must be false when configuring access to individual sub-resources. By
	// default, the value is true. See Configuration Advanced Options
	// (http://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/es-createupdatedomains.html#es-createdomain-configure-advanced-options)
	// for more information.
	AdvancedOptions map[string]*string

	// Specifies advanced security options.
	AdvancedSecurityOptions *types.AdvancedSecurityOptionsInput

	// Options to specify the Cognito user and identity pools for Kibana
	// authentication. For more information, see Amazon Cognito Authentication for
	// Kibana
	// (http://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/es-cognito-auth.html).
	CognitoOptions *types.CognitoOptions

	// Options to specify configuration that will be applied to the domain endpoint.
	DomainEndpointOptions *types.DomainEndpointOptions

	// Specify the type and size of the EBS volume that you want to use.
	EBSOptions *types.EBSOptions

	// The type and number of instances to instantiate for the domain cluster.
	ElasticsearchClusterConfig *types.ElasticsearchClusterConfig

	// Map of LogType and LogPublishingOption, each containing options to publish a
	// given type of Elasticsearch log.
	LogPublishingOptions map[string]*types.LogPublishingOption

	// Option to set the time, in UTC format, for the daily automated snapshot. Default
	// value is 0 hours.
	SnapshotOptions *types.SnapshotOptions

	// Options to specify the subnets and security groups for VPC endpoint. For more
	// information, see Creating a VPC
	// (http://docs.aws.amazon.com/elasticsearch-service/latest/developerguide/es-vpc.html#es-creating-vpc)
	// in VPC Endpoints for Amazon Elasticsearch Service Domains
	VPCOptions *types.VPCOptions
}

// The result of an UpdateElasticsearchDomain request. Contains the status of the
// Elasticsearch domain being updated.
type UpdateElasticsearchDomainConfigOutput struct {

	// The status of the updated Elasticsearch domain.
	//
	// This member is required.
	DomainConfig *types.ElasticsearchDomainConfig

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationUpdateElasticsearchDomainConfigMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateElasticsearchDomainConfig{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateElasticsearchDomainConfig{}, middleware.After)
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
	addOpUpdateElasticsearchDomainConfigValidationMiddleware(stack)
	stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateElasticsearchDomainConfig(options.Region), middleware.Before)
	addRequestIDRetrieverMiddleware(stack)
	addResponseErrorMiddleware(stack)
	return nil
}

func newServiceMetadataMiddleware_opUpdateElasticsearchDomainConfig(region string) awsmiddleware.RegisterServiceMetadata {
	return awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "es",
		OperationName: "UpdateElasticsearchDomainConfig",
	}
}
