// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mediapackage

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
	"github.com/aws/aws-sdk-go-v2/private/protocol"
)

// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediapackage-2017-10-12/DescribeOriginEndpointRequest
type DescribeOriginEndpointInput struct {
	_ struct{} `type:"structure"`

	// Id is a required field
	Id *string `location:"uri" locationName:"id" type:"string" required:"true"`
}

// String returns the string representation
func (s DescribeOriginEndpointInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeOriginEndpointInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "DescribeOriginEndpointInput"}

	if s.Id == nil {
		invalidParams.Add(aws.NewErrParamRequired("Id"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeOriginEndpointInput) MarshalFields(e protocol.FieldEncoder) error {
	e.SetValue(protocol.HeaderTarget, "Content-Type", protocol.StringValue("application/x-amz-json-1.1"), protocol.Metadata{})

	if s.Id != nil {
		v := *s.Id

		metadata := protocol.Metadata{}
		e.SetValue(protocol.PathTarget, "id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	return nil
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediapackage-2017-10-12/DescribeOriginEndpointResponse
type DescribeOriginEndpointOutput struct {
	_ struct{} `type:"structure"`

	Arn *string `locationName:"arn" type:"string"`

	ChannelId *string `locationName:"channelId" type:"string"`

	// A Common Media Application Format (CMAF) packaging configuration.
	CmafPackage *CmafPackage `locationName:"cmafPackage" type:"structure"`

	// A Dynamic Adaptive Streaming over HTTP (DASH) packaging configuration.
	DashPackage *DashPackage `locationName:"dashPackage" type:"structure"`

	Description *string `locationName:"description" type:"string"`

	// An HTTP Live Streaming (HLS) packaging configuration.
	HlsPackage *HlsPackage `locationName:"hlsPackage" type:"structure"`

	Id *string `locationName:"id" type:"string"`

	ManifestName *string `locationName:"manifestName" type:"string"`

	// A Microsoft Smooth Streaming (MSS) packaging configuration.
	MssPackage *MssPackage `locationName:"mssPackage" type:"structure"`

	StartoverWindowSeconds *int64 `locationName:"startoverWindowSeconds" type:"integer"`

	// A collection of tags associated with a resource
	Tags map[string]string `locationName:"tags" type:"map"`

	TimeDelaySeconds *int64 `locationName:"timeDelaySeconds" type:"integer"`

	Url *string `locationName:"url" type:"string"`

	Whitelist []string `locationName:"whitelist" type:"list"`
}

// String returns the string representation
func (s DescribeOriginEndpointOutput) String() string {
	return awsutil.Prettify(s)
}

// MarshalFields encodes the AWS API shape using the passed in protocol encoder.
func (s DescribeOriginEndpointOutput) MarshalFields(e protocol.FieldEncoder) error {
	if s.Arn != nil {
		v := *s.Arn

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "arn", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ChannelId != nil {
		v := *s.ChannelId

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "channelId", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.CmafPackage != nil {
		v := s.CmafPackage

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "cmafPackage", v, metadata)
	}
	if s.DashPackage != nil {
		v := s.DashPackage

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "dashPackage", v, metadata)
	}
	if s.Description != nil {
		v := *s.Description

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "description", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.HlsPackage != nil {
		v := s.HlsPackage

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "hlsPackage", v, metadata)
	}
	if s.Id != nil {
		v := *s.Id

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "id", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.ManifestName != nil {
		v := *s.ManifestName

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "manifestName", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if s.MssPackage != nil {
		v := s.MssPackage

		metadata := protocol.Metadata{}
		e.SetFields(protocol.BodyTarget, "mssPackage", v, metadata)
	}
	if s.StartoverWindowSeconds != nil {
		v := *s.StartoverWindowSeconds

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "startoverWindowSeconds", protocol.Int64Value(v), metadata)
	}
	if len(s.Tags) > 0 {
		v := s.Tags

		metadata := protocol.Metadata{}
		ms0 := e.Map(protocol.BodyTarget, "tags", metadata)
		ms0.Start()
		for k1, v1 := range v {
			ms0.MapSetValue(k1, protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ms0.End()

	}
	if s.TimeDelaySeconds != nil {
		v := *s.TimeDelaySeconds

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "timeDelaySeconds", protocol.Int64Value(v), metadata)
	}
	if s.Url != nil {
		v := *s.Url

		metadata := protocol.Metadata{}
		e.SetValue(protocol.BodyTarget, "url", protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v)}, metadata)
	}
	if len(s.Whitelist) > 0 {
		v := s.Whitelist

		metadata := protocol.Metadata{}
		ls0 := e.List(protocol.BodyTarget, "whitelist", metadata)
		ls0.Start()
		for _, v1 := range v {
			ls0.ListAddValue(protocol.QuotedValue{ValueMarshaler: protocol.StringValue(v1)})
		}
		ls0.End()

	}
	return nil
}

const opDescribeOriginEndpoint = "DescribeOriginEndpoint"

// DescribeOriginEndpointRequest returns a request value for making API operation for
// AWS Elemental MediaPackage.
//
// Gets details about an existing OriginEndpoint.
//
//    // Example sending a request using DescribeOriginEndpointRequest.
//    req := client.DescribeOriginEndpointRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/mediapackage-2017-10-12/DescribeOriginEndpoint
func (c *Client) DescribeOriginEndpointRequest(input *DescribeOriginEndpointInput) DescribeOriginEndpointRequest {
	op := &aws.Operation{
		Name:       opDescribeOriginEndpoint,
		HTTPMethod: "GET",
		HTTPPath:   "/origin_endpoints/{id}",
	}

	if input == nil {
		input = &DescribeOriginEndpointInput{}
	}

	req := c.newRequest(op, input, &DescribeOriginEndpointOutput{})
	return DescribeOriginEndpointRequest{Request: req, Input: input, Copy: c.DescribeOriginEndpointRequest}
}

// DescribeOriginEndpointRequest is the request type for the
// DescribeOriginEndpoint API operation.
type DescribeOriginEndpointRequest struct {
	*aws.Request
	Input *DescribeOriginEndpointInput
	Copy  func(*DescribeOriginEndpointInput) DescribeOriginEndpointRequest
}

// Send marshals and sends the DescribeOriginEndpoint API request.
func (r DescribeOriginEndpointRequest) Send(ctx context.Context) (*DescribeOriginEndpointResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &DescribeOriginEndpointResponse{
		DescribeOriginEndpointOutput: r.Request.Data.(*DescribeOriginEndpointOutput),
		response:                     &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// DescribeOriginEndpointResponse is the response type for the
// DescribeOriginEndpoint API operation.
type DescribeOriginEndpointResponse struct {
	*DescribeOriginEndpointOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// DescribeOriginEndpoint request.
func (r *DescribeOriginEndpointResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}