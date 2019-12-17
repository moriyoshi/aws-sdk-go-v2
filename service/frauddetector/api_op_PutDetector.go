// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package frauddetector

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type PutDetectorInput struct {
	_ struct{} `type:"structure"`

	// The description of the detector.
	Description *string `locationName:"description" min:"1" type:"string"`

	// The detector ID.
	//
	// DetectorId is a required field
	DetectorId *string `locationName:"detectorId" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s PutDetectorInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PutDetectorInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "PutDetectorInput"}
	if s.Description != nil && len(*s.Description) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("Description", 1))
	}

	if s.DetectorId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DetectorId"))
	}
	if s.DetectorId != nil && len(*s.DetectorId) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("DetectorId", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type PutDetectorOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s PutDetectorOutput) String() string {
	return awsutil.Prettify(s)
}

const opPutDetector = "PutDetector"

// PutDetectorRequest returns a request value for making API operation for
// Amazon Fraud Detector.
//
// Creates or updates a detector.
//
//    // Example sending a request using PutDetectorRequest.
//    req := client.PutDetectorRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/frauddetector-2019-11-15/PutDetector
func (c *Client) PutDetectorRequest(input *PutDetectorInput) PutDetectorRequest {
	op := &aws.Operation{
		Name:       opPutDetector,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &PutDetectorInput{}
	}

	req := c.newRequest(op, input, &PutDetectorOutput{})
	return PutDetectorRequest{Request: req, Input: input, Copy: c.PutDetectorRequest}
}

// PutDetectorRequest is the request type for the
// PutDetector API operation.
type PutDetectorRequest struct {
	*aws.Request
	Input *PutDetectorInput
	Copy  func(*PutDetectorInput) PutDetectorRequest
}

// Send marshals and sends the PutDetector API request.
func (r PutDetectorRequest) Send(ctx context.Context) (*PutDetectorResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &PutDetectorResponse{
		PutDetectorOutput: r.Request.Data.(*PutDetectorOutput),
		response:          &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// PutDetectorResponse is the response type for the
// PutDetector API operation.
type PutDetectorResponse struct {
	*PutDetectorOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// PutDetector request.
func (r *PutDetectorResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}