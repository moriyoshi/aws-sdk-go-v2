// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package directoryservice

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

// Contains the inputs for the CreateComputer operation.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ds-2015-04-16/CreateComputerRequest
type CreateComputerInput struct {
	_ struct{} `type:"structure"`

	// An array of Attribute objects that contain any LDAP attributes to apply to
	// the computer account.
	ComputerAttributes []Attribute `type:"list"`

	// The name of the computer account.
	//
	// ComputerName is a required field
	ComputerName *string `min:"1" type:"string" required:"true"`

	// The identifier of the directory in which to create the computer account.
	//
	// DirectoryId is a required field
	DirectoryId *string `type:"string" required:"true"`

	// The fully-qualified distinguished name of the organizational unit to place
	// the computer account in.
	OrganizationalUnitDistinguishedName *string `min:"1" type:"string"`

	// A one-time password that is used to join the computer to the directory. You
	// should generate a random, strong password to use for this parameter.
	//
	// Password is a required field
	Password *string `min:"8" type:"string" required:"true"`
}

// String returns the string representation
func (s CreateComputerInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateComputerInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "CreateComputerInput"}

	if s.ComputerName == nil {
		invalidParams.Add(aws.NewErrParamRequired("ComputerName"))
	}
	if s.ComputerName != nil && len(*s.ComputerName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("ComputerName", 1))
	}

	if s.DirectoryId == nil {
		invalidParams.Add(aws.NewErrParamRequired("DirectoryId"))
	}
	if s.OrganizationalUnitDistinguishedName != nil && len(*s.OrganizationalUnitDistinguishedName) < 1 {
		invalidParams.Add(aws.NewErrParamMinLen("OrganizationalUnitDistinguishedName", 1))
	}

	if s.Password == nil {
		invalidParams.Add(aws.NewErrParamRequired("Password"))
	}
	if s.Password != nil && len(*s.Password) < 8 {
		invalidParams.Add(aws.NewErrParamMinLen("Password", 8))
	}
	if s.ComputerAttributes != nil {
		for i, v := range s.ComputerAttributes {
			if err := v.Validate(); err != nil {
				invalidParams.AddNested(fmt.Sprintf("%s[%v]", "ComputerAttributes", i), err.(aws.ErrInvalidParams))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// Contains the results for the CreateComputer operation.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ds-2015-04-16/CreateComputerResult
type CreateComputerOutput struct {
	_ struct{} `type:"structure"`

	// A Computer object that represents the computer account.
	Computer *Computer `type:"structure"`
}

// String returns the string representation
func (s CreateComputerOutput) String() string {
	return awsutil.Prettify(s)
}

const opCreateComputer = "CreateComputer"

// CreateComputerRequest returns a request value for making API operation for
// AWS Directory Service.
//
// Creates a computer account in the specified directory, and joins the computer
// to the directory.
//
//    // Example sending a request using CreateComputerRequest.
//    req := client.CreateComputerRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ds-2015-04-16/CreateComputer
func (c *Client) CreateComputerRequest(input *CreateComputerInput) CreateComputerRequest {
	op := &aws.Operation{
		Name:       opCreateComputer,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateComputerInput{}
	}

	req := c.newRequest(op, input, &CreateComputerOutput{})
	return CreateComputerRequest{Request: req, Input: input, Copy: c.CreateComputerRequest}
}

// CreateComputerRequest is the request type for the
// CreateComputer API operation.
type CreateComputerRequest struct {
	*aws.Request
	Input *CreateComputerInput
	Copy  func(*CreateComputerInput) CreateComputerRequest
}

// Send marshals and sends the CreateComputer API request.
func (r CreateComputerRequest) Send(ctx context.Context) (*CreateComputerResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &CreateComputerResponse{
		CreateComputerOutput: r.Request.Data.(*CreateComputerOutput),
		response:             &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// CreateComputerResponse is the response type for the
// CreateComputer API operation.
type CreateComputerResponse struct {
	*CreateComputerOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// CreateComputer request.
func (r *CreateComputerResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}