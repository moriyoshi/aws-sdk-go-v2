// Code generated by smithy-go-codegen DO NOT EDIT.

package sso

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws/protocol/restjson"
	"github.com/aws/aws-sdk-go-v2/service/sso/types"
	smithy "github.com/awslabs/smithy-go"
	smithyio "github.com/awslabs/smithy-go/io"
	smithyjson "github.com/awslabs/smithy-go/json"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
	"io"
	"strings"
)

type awsRestjson1_deserializeOpGetRoleCredentials struct {
}

func (*awsRestjson1_deserializeOpGetRoleCredentials) ID() string {
	return "OperationDeserializer"
}

func (m *awsRestjson1_deserializeOpGetRoleCredentials) HandleDeserialize(ctx context.Context, in middleware.DeserializeInput, next middleware.DeserializeHandler) (
	out middleware.DeserializeOutput, metadata middleware.Metadata, err error,
) {
	out, metadata, err = next.HandleDeserialize(ctx, in)
	if err != nil {
		return out, metadata, err
	}

	response, ok := out.RawResponse.(*smithyhttp.Response)
	if !ok {
		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
	}

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return out, metadata, awsRestjson1_deserializeOpErrorGetRoleCredentials(response)
	}
	output := &GetRoleCredentialsOutput{}
	out.Result = output

	buff := make([]byte, 1024)
	ringBuffer := smithyio.NewRingBuffer(buff)

	body := io.TeeReader(response.Body, ringBuffer)

	decoder := json.NewDecoder(body)
	decoder.UseNumber()

	err = awsRestjson1_deserializeOpDocumentGetRoleCredentialsOutput(&output, decoder)
	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		return out, metadata, &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body with invalid JSON, %w", err),
			Snapshot: snapshot.Bytes(),
		}
	}

	return out, metadata, err
}

func awsRestjson1_deserializeOpErrorGetRoleCredentials(response *smithyhttp.Response) error {
	var errorBuffer bytes.Buffer
	if _, err := io.Copy(&errorBuffer, response.Body); err != nil {
		return &smithy.DeserializationError{Err: fmt.Errorf("failed to copy error response body, %w", err)}
	}
	errorBody := bytes.NewReader(errorBuffer.Bytes())

	errorCode := "UnknownError"
	errorMessage := errorCode

	code := response.Header.Get("X-Amzn-ErrorType")
	if len(code) != 0 {
		errorCode = restjson.SanitizeErrorCode(code)
	}

	buff := make([]byte, 1024)
	ringBuffer := smithyio.NewRingBuffer(buff)

	body := io.TeeReader(errorBody, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()

	code, message, err := restjson.GetErrorInfo(decoder)
	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		return &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
	}

	errorBody.Seek(0, io.SeekStart)
	if len(code) != 0 {
		errorCode = restjson.SanitizeErrorCode(code)
	}
	if len(message) != 0 {
		errorMessage = message
	}

	switch {
	case strings.EqualFold("InvalidRequestException", errorCode):
		return awsRestjson1_deserializeErrorInvalidRequestException(response, errorBody)

	case strings.EqualFold("ResourceNotFoundException", errorCode):
		return awsRestjson1_deserializeErrorResourceNotFoundException(response, errorBody)

	case strings.EqualFold("TooManyRequestsException", errorCode):
		return awsRestjson1_deserializeErrorTooManyRequestsException(response, errorBody)

	case strings.EqualFold("UnauthorizedException", errorCode):
		return awsRestjson1_deserializeErrorUnauthorizedException(response, errorBody)

	default:
		genericError := &smithy.GenericAPIError{
			Code:    errorCode,
			Message: errorMessage,
		}
		return genericError

	}
}

func awsRestjson1_deserializeOpDocumentGetRoleCredentialsOutput(v **GetRoleCredentialsOutput, decoder *json.Decoder) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	startToken, err := decoder.Token()
	if err == io.EOF {
		return nil
	}
	if err != nil {
		return err
	}
	if startToken == nil {
		return nil
	}
	if t, ok := startToken.(json.Delim); !ok || t != '{' {
		return fmt.Errorf("expect `{` as start token")
	}

	var sv *GetRoleCredentialsOutput
	if *v == nil {
		sv = &GetRoleCredentialsOutput{}
	} else {
		sv = *v
	}

	for decoder.More() {
		t, err := decoder.Token()
		if err != nil {
			return err
		}
		switch t {
		case "roleCredentials":
			if err := awsRestjson1_deserializeDocumentRoleCredentials(&sv.RoleCredentials, decoder); err != nil {
				return err
			}

		default:
			err := smithyjson.DiscardUnknownField(decoder)
			if err != nil {
				return err
			}

		}
	}
	endToken, err := decoder.Token()
	if err != nil {
		return err
	}
	if t, ok := endToken.(json.Delim); !ok || t != '}' {
		return fmt.Errorf("expect `}` as end token")
	}

	*v = sv
	return nil
}

type awsRestjson1_deserializeOpListAccountRoles struct {
}

func (*awsRestjson1_deserializeOpListAccountRoles) ID() string {
	return "OperationDeserializer"
}

func (m *awsRestjson1_deserializeOpListAccountRoles) HandleDeserialize(ctx context.Context, in middleware.DeserializeInput, next middleware.DeserializeHandler) (
	out middleware.DeserializeOutput, metadata middleware.Metadata, err error,
) {
	out, metadata, err = next.HandleDeserialize(ctx, in)
	if err != nil {
		return out, metadata, err
	}

	response, ok := out.RawResponse.(*smithyhttp.Response)
	if !ok {
		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
	}

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return out, metadata, awsRestjson1_deserializeOpErrorListAccountRoles(response)
	}
	output := &ListAccountRolesOutput{}
	out.Result = output

	buff := make([]byte, 1024)
	ringBuffer := smithyio.NewRingBuffer(buff)

	body := io.TeeReader(response.Body, ringBuffer)

	decoder := json.NewDecoder(body)
	decoder.UseNumber()

	err = awsRestjson1_deserializeOpDocumentListAccountRolesOutput(&output, decoder)
	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		return out, metadata, &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body with invalid JSON, %w", err),
			Snapshot: snapshot.Bytes(),
		}
	}

	return out, metadata, err
}

func awsRestjson1_deserializeOpErrorListAccountRoles(response *smithyhttp.Response) error {
	var errorBuffer bytes.Buffer
	if _, err := io.Copy(&errorBuffer, response.Body); err != nil {
		return &smithy.DeserializationError{Err: fmt.Errorf("failed to copy error response body, %w", err)}
	}
	errorBody := bytes.NewReader(errorBuffer.Bytes())

	errorCode := "UnknownError"
	errorMessage := errorCode

	code := response.Header.Get("X-Amzn-ErrorType")
	if len(code) != 0 {
		errorCode = restjson.SanitizeErrorCode(code)
	}

	buff := make([]byte, 1024)
	ringBuffer := smithyio.NewRingBuffer(buff)

	body := io.TeeReader(errorBody, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()

	code, message, err := restjson.GetErrorInfo(decoder)
	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		return &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
	}

	errorBody.Seek(0, io.SeekStart)
	if len(code) != 0 {
		errorCode = restjson.SanitizeErrorCode(code)
	}
	if len(message) != 0 {
		errorMessage = message
	}

	switch {
	case strings.EqualFold("InvalidRequestException", errorCode):
		return awsRestjson1_deserializeErrorInvalidRequestException(response, errorBody)

	case strings.EqualFold("ResourceNotFoundException", errorCode):
		return awsRestjson1_deserializeErrorResourceNotFoundException(response, errorBody)

	case strings.EqualFold("TooManyRequestsException", errorCode):
		return awsRestjson1_deserializeErrorTooManyRequestsException(response, errorBody)

	case strings.EqualFold("UnauthorizedException", errorCode):
		return awsRestjson1_deserializeErrorUnauthorizedException(response, errorBody)

	default:
		genericError := &smithy.GenericAPIError{
			Code:    errorCode,
			Message: errorMessage,
		}
		return genericError

	}
}

func awsRestjson1_deserializeOpDocumentListAccountRolesOutput(v **ListAccountRolesOutput, decoder *json.Decoder) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	startToken, err := decoder.Token()
	if err == io.EOF {
		return nil
	}
	if err != nil {
		return err
	}
	if startToken == nil {
		return nil
	}
	if t, ok := startToken.(json.Delim); !ok || t != '{' {
		return fmt.Errorf("expect `{` as start token")
	}

	var sv *ListAccountRolesOutput
	if *v == nil {
		sv = &ListAccountRolesOutput{}
	} else {
		sv = *v
	}

	for decoder.More() {
		t, err := decoder.Token()
		if err != nil {
			return err
		}
		switch t {
		case "nextToken":
			val, err := decoder.Token()
			if err != nil {
				return err
			}
			if val != nil {
				jtv, ok := val.(string)
				if !ok {
					return fmt.Errorf("expected NextTokenType to be of type string, got %T instead", val)
				}
				sv.NextToken = &jtv
			}

		case "roleList":
			if err := awsRestjson1_deserializeDocumentRoleListType(&sv.RoleList, decoder); err != nil {
				return err
			}

		default:
			err := smithyjson.DiscardUnknownField(decoder)
			if err != nil {
				return err
			}

		}
	}
	endToken, err := decoder.Token()
	if err != nil {
		return err
	}
	if t, ok := endToken.(json.Delim); !ok || t != '}' {
		return fmt.Errorf("expect `}` as end token")
	}

	*v = sv
	return nil
}

type awsRestjson1_deserializeOpListAccounts struct {
}

func (*awsRestjson1_deserializeOpListAccounts) ID() string {
	return "OperationDeserializer"
}

func (m *awsRestjson1_deserializeOpListAccounts) HandleDeserialize(ctx context.Context, in middleware.DeserializeInput, next middleware.DeserializeHandler) (
	out middleware.DeserializeOutput, metadata middleware.Metadata, err error,
) {
	out, metadata, err = next.HandleDeserialize(ctx, in)
	if err != nil {
		return out, metadata, err
	}

	response, ok := out.RawResponse.(*smithyhttp.Response)
	if !ok {
		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
	}

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return out, metadata, awsRestjson1_deserializeOpErrorListAccounts(response)
	}
	output := &ListAccountsOutput{}
	out.Result = output

	buff := make([]byte, 1024)
	ringBuffer := smithyio.NewRingBuffer(buff)

	body := io.TeeReader(response.Body, ringBuffer)

	decoder := json.NewDecoder(body)
	decoder.UseNumber()

	err = awsRestjson1_deserializeOpDocumentListAccountsOutput(&output, decoder)
	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		return out, metadata, &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body with invalid JSON, %w", err),
			Snapshot: snapshot.Bytes(),
		}
	}

	return out, metadata, err
}

func awsRestjson1_deserializeOpErrorListAccounts(response *smithyhttp.Response) error {
	var errorBuffer bytes.Buffer
	if _, err := io.Copy(&errorBuffer, response.Body); err != nil {
		return &smithy.DeserializationError{Err: fmt.Errorf("failed to copy error response body, %w", err)}
	}
	errorBody := bytes.NewReader(errorBuffer.Bytes())

	errorCode := "UnknownError"
	errorMessage := errorCode

	code := response.Header.Get("X-Amzn-ErrorType")
	if len(code) != 0 {
		errorCode = restjson.SanitizeErrorCode(code)
	}

	buff := make([]byte, 1024)
	ringBuffer := smithyio.NewRingBuffer(buff)

	body := io.TeeReader(errorBody, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()

	code, message, err := restjson.GetErrorInfo(decoder)
	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		return &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
	}

	errorBody.Seek(0, io.SeekStart)
	if len(code) != 0 {
		errorCode = restjson.SanitizeErrorCode(code)
	}
	if len(message) != 0 {
		errorMessage = message
	}

	switch {
	case strings.EqualFold("InvalidRequestException", errorCode):
		return awsRestjson1_deserializeErrorInvalidRequestException(response, errorBody)

	case strings.EqualFold("ResourceNotFoundException", errorCode):
		return awsRestjson1_deserializeErrorResourceNotFoundException(response, errorBody)

	case strings.EqualFold("TooManyRequestsException", errorCode):
		return awsRestjson1_deserializeErrorTooManyRequestsException(response, errorBody)

	case strings.EqualFold("UnauthorizedException", errorCode):
		return awsRestjson1_deserializeErrorUnauthorizedException(response, errorBody)

	default:
		genericError := &smithy.GenericAPIError{
			Code:    errorCode,
			Message: errorMessage,
		}
		return genericError

	}
}

func awsRestjson1_deserializeOpDocumentListAccountsOutput(v **ListAccountsOutput, decoder *json.Decoder) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	startToken, err := decoder.Token()
	if err == io.EOF {
		return nil
	}
	if err != nil {
		return err
	}
	if startToken == nil {
		return nil
	}
	if t, ok := startToken.(json.Delim); !ok || t != '{' {
		return fmt.Errorf("expect `{` as start token")
	}

	var sv *ListAccountsOutput
	if *v == nil {
		sv = &ListAccountsOutput{}
	} else {
		sv = *v
	}

	for decoder.More() {
		t, err := decoder.Token()
		if err != nil {
			return err
		}
		switch t {
		case "accountList":
			if err := awsRestjson1_deserializeDocumentAccountListType(&sv.AccountList, decoder); err != nil {
				return err
			}

		case "nextToken":
			val, err := decoder.Token()
			if err != nil {
				return err
			}
			if val != nil {
				jtv, ok := val.(string)
				if !ok {
					return fmt.Errorf("expected NextTokenType to be of type string, got %T instead", val)
				}
				sv.NextToken = &jtv
			}

		default:
			err := smithyjson.DiscardUnknownField(decoder)
			if err != nil {
				return err
			}

		}
	}
	endToken, err := decoder.Token()
	if err != nil {
		return err
	}
	if t, ok := endToken.(json.Delim); !ok || t != '}' {
		return fmt.Errorf("expect `}` as end token")
	}

	*v = sv
	return nil
}

type awsRestjson1_deserializeOpLogout struct {
}

func (*awsRestjson1_deserializeOpLogout) ID() string {
	return "OperationDeserializer"
}

func (m *awsRestjson1_deserializeOpLogout) HandleDeserialize(ctx context.Context, in middleware.DeserializeInput, next middleware.DeserializeHandler) (
	out middleware.DeserializeOutput, metadata middleware.Metadata, err error,
) {
	out, metadata, err = next.HandleDeserialize(ctx, in)
	if err != nil {
		return out, metadata, err
	}

	response, ok := out.RawResponse.(*smithyhttp.Response)
	if !ok {
		return out, metadata, &smithy.DeserializationError{Err: fmt.Errorf("unknown transport type %T", out.RawResponse)}
	}

	if response.StatusCode < 200 || response.StatusCode >= 300 {
		return out, metadata, awsRestjson1_deserializeOpErrorLogout(response)
	}
	output := &LogoutOutput{}
	out.Result = output

	return out, metadata, err
}

func awsRestjson1_deserializeOpErrorLogout(response *smithyhttp.Response) error {
	var errorBuffer bytes.Buffer
	if _, err := io.Copy(&errorBuffer, response.Body); err != nil {
		return &smithy.DeserializationError{Err: fmt.Errorf("failed to copy error response body, %w", err)}
	}
	errorBody := bytes.NewReader(errorBuffer.Bytes())

	errorCode := "UnknownError"
	errorMessage := errorCode

	code := response.Header.Get("X-Amzn-ErrorType")
	if len(code) != 0 {
		errorCode = restjson.SanitizeErrorCode(code)
	}

	buff := make([]byte, 1024)
	ringBuffer := smithyio.NewRingBuffer(buff)

	body := io.TeeReader(errorBody, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()

	code, message, err := restjson.GetErrorInfo(decoder)
	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		return &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
	}

	errorBody.Seek(0, io.SeekStart)
	if len(code) != 0 {
		errorCode = restjson.SanitizeErrorCode(code)
	}
	if len(message) != 0 {
		errorMessage = message
	}

	switch {
	case strings.EqualFold("InvalidRequestException", errorCode):
		return awsRestjson1_deserializeErrorInvalidRequestException(response, errorBody)

	case strings.EqualFold("TooManyRequestsException", errorCode):
		return awsRestjson1_deserializeErrorTooManyRequestsException(response, errorBody)

	case strings.EqualFold("UnauthorizedException", errorCode):
		return awsRestjson1_deserializeErrorUnauthorizedException(response, errorBody)

	default:
		genericError := &smithy.GenericAPIError{
			Code:    errorCode,
			Message: errorMessage,
		}
		return genericError

	}
}

func awsRestjson1_deserializeErrorInvalidRequestException(response *smithyhttp.Response, errorBody *bytes.Reader) error {
	output := &types.InvalidRequestException{}
	buff := make([]byte, 1024)
	ringBuffer := smithyio.NewRingBuffer(buff)

	body := io.TeeReader(errorBody, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()

	err := awsRestjson1_deserializeDocumentInvalidRequestException(&output, decoder)

	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		return &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
	}

	errorBody.Seek(0, io.SeekStart)

	return output
}

func awsRestjson1_deserializeErrorResourceNotFoundException(response *smithyhttp.Response, errorBody *bytes.Reader) error {
	output := &types.ResourceNotFoundException{}
	buff := make([]byte, 1024)
	ringBuffer := smithyio.NewRingBuffer(buff)

	body := io.TeeReader(errorBody, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()

	err := awsRestjson1_deserializeDocumentResourceNotFoundException(&output, decoder)

	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		return &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
	}

	errorBody.Seek(0, io.SeekStart)

	return output
}

func awsRestjson1_deserializeErrorTooManyRequestsException(response *smithyhttp.Response, errorBody *bytes.Reader) error {
	output := &types.TooManyRequestsException{}
	buff := make([]byte, 1024)
	ringBuffer := smithyio.NewRingBuffer(buff)

	body := io.TeeReader(errorBody, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()

	err := awsRestjson1_deserializeDocumentTooManyRequestsException(&output, decoder)

	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		return &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
	}

	errorBody.Seek(0, io.SeekStart)

	return output
}

func awsRestjson1_deserializeErrorUnauthorizedException(response *smithyhttp.Response, errorBody *bytes.Reader) error {
	output := &types.UnauthorizedException{}
	buff := make([]byte, 1024)
	ringBuffer := smithyio.NewRingBuffer(buff)

	body := io.TeeReader(errorBody, ringBuffer)
	decoder := json.NewDecoder(body)
	decoder.UseNumber()

	err := awsRestjson1_deserializeDocumentUnauthorizedException(&output, decoder)

	if err != nil {
		var snapshot bytes.Buffer
		io.Copy(&snapshot, ringBuffer)
		return &smithy.DeserializationError{
			Err:      fmt.Errorf("failed to decode response body, %w", err),
			Snapshot: snapshot.Bytes(),
		}
	}

	errorBody.Seek(0, io.SeekStart)

	return output
}

func awsRestjson1_deserializeDocumentAccountInfo(v **types.AccountInfo, decoder *json.Decoder) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	startToken, err := decoder.Token()
	if err == io.EOF {
		return nil
	}
	if err != nil {
		return err
	}
	if startToken == nil {
		return nil
	}
	if t, ok := startToken.(json.Delim); !ok || t != '{' {
		return fmt.Errorf("expect `{` as start token")
	}

	var sv *types.AccountInfo
	if *v == nil {
		sv = &types.AccountInfo{}
	} else {
		sv = *v
	}

	for decoder.More() {
		t, err := decoder.Token()
		if err != nil {
			return err
		}
		switch t {
		case "accountId":
			val, err := decoder.Token()
			if err != nil {
				return err
			}
			if val != nil {
				jtv, ok := val.(string)
				if !ok {
					return fmt.Errorf("expected AccountIdType to be of type string, got %T instead", val)
				}
				sv.AccountId = &jtv
			}

		case "accountName":
			val, err := decoder.Token()
			if err != nil {
				return err
			}
			if val != nil {
				jtv, ok := val.(string)
				if !ok {
					return fmt.Errorf("expected AccountNameType to be of type string, got %T instead", val)
				}
				sv.AccountName = &jtv
			}

		case "emailAddress":
			val, err := decoder.Token()
			if err != nil {
				return err
			}
			if val != nil {
				jtv, ok := val.(string)
				if !ok {
					return fmt.Errorf("expected EmailAddressType to be of type string, got %T instead", val)
				}
				sv.EmailAddress = &jtv
			}

		default:
			err := smithyjson.DiscardUnknownField(decoder)
			if err != nil {
				return err
			}

		}
	}
	endToken, err := decoder.Token()
	if err != nil {
		return err
	}
	if t, ok := endToken.(json.Delim); !ok || t != '}' {
		return fmt.Errorf("expect `}` as end token")
	}

	*v = sv
	return nil
}

func awsRestjson1_deserializeDocumentAccountListType(v *[]*types.AccountInfo, decoder *json.Decoder) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	startToken, err := decoder.Token()
	if err == io.EOF {
		return nil
	}
	if err != nil {
		return err
	}
	if startToken == nil {
		return nil
	}
	if t, ok := startToken.(json.Delim); !ok || t != '[' {
		return fmt.Errorf("expect `[` as start token")
	}

	var cv []*types.AccountInfo
	if *v == nil {
		cv = []*types.AccountInfo{}
	} else {
		cv = *v
	}

	for decoder.More() {
		var col *types.AccountInfo
		if err := awsRestjson1_deserializeDocumentAccountInfo(&col, decoder); err != nil {
			return err
		}
		cv = append(cv, col)

	}
	endToken, err := decoder.Token()
	if err != nil {
		return err
	}
	if t, ok := endToken.(json.Delim); !ok || t != ']' {
		return fmt.Errorf("expect `]` as end token")
	}

	*v = cv
	return nil
}

func awsRestjson1_deserializeDocumentInvalidRequestException(v **types.InvalidRequestException, decoder *json.Decoder) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	startToken, err := decoder.Token()
	if err == io.EOF {
		return nil
	}
	if err != nil {
		return err
	}
	if startToken == nil {
		return nil
	}
	if t, ok := startToken.(json.Delim); !ok || t != '{' {
		return fmt.Errorf("expect `{` as start token")
	}

	var sv *types.InvalidRequestException
	if *v == nil {
		sv = &types.InvalidRequestException{}
	} else {
		sv = *v
	}

	for decoder.More() {
		t, err := decoder.Token()
		if err != nil {
			return err
		}
		switch t {
		case "message":
			val, err := decoder.Token()
			if err != nil {
				return err
			}
			if val != nil {
				jtv, ok := val.(string)
				if !ok {
					return fmt.Errorf("expected ErrorDescription to be of type string, got %T instead", val)
				}
				sv.Message = &jtv
			}

		default:
			err := smithyjson.DiscardUnknownField(decoder)
			if err != nil {
				return err
			}

		}
	}
	endToken, err := decoder.Token()
	if err != nil {
		return err
	}
	if t, ok := endToken.(json.Delim); !ok || t != '}' {
		return fmt.Errorf("expect `}` as end token")
	}

	*v = sv
	return nil
}

func awsRestjson1_deserializeDocumentResourceNotFoundException(v **types.ResourceNotFoundException, decoder *json.Decoder) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	startToken, err := decoder.Token()
	if err == io.EOF {
		return nil
	}
	if err != nil {
		return err
	}
	if startToken == nil {
		return nil
	}
	if t, ok := startToken.(json.Delim); !ok || t != '{' {
		return fmt.Errorf("expect `{` as start token")
	}

	var sv *types.ResourceNotFoundException
	if *v == nil {
		sv = &types.ResourceNotFoundException{}
	} else {
		sv = *v
	}

	for decoder.More() {
		t, err := decoder.Token()
		if err != nil {
			return err
		}
		switch t {
		case "message":
			val, err := decoder.Token()
			if err != nil {
				return err
			}
			if val != nil {
				jtv, ok := val.(string)
				if !ok {
					return fmt.Errorf("expected ErrorDescription to be of type string, got %T instead", val)
				}
				sv.Message = &jtv
			}

		default:
			err := smithyjson.DiscardUnknownField(decoder)
			if err != nil {
				return err
			}

		}
	}
	endToken, err := decoder.Token()
	if err != nil {
		return err
	}
	if t, ok := endToken.(json.Delim); !ok || t != '}' {
		return fmt.Errorf("expect `}` as end token")
	}

	*v = sv
	return nil
}

func awsRestjson1_deserializeDocumentRoleCredentials(v **types.RoleCredentials, decoder *json.Decoder) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	startToken, err := decoder.Token()
	if err == io.EOF {
		return nil
	}
	if err != nil {
		return err
	}
	if startToken == nil {
		return nil
	}
	if t, ok := startToken.(json.Delim); !ok || t != '{' {
		return fmt.Errorf("expect `{` as start token")
	}

	var sv *types.RoleCredentials
	if *v == nil {
		sv = &types.RoleCredentials{}
	} else {
		sv = *v
	}

	for decoder.More() {
		t, err := decoder.Token()
		if err != nil {
			return err
		}
		switch t {
		case "accessKeyId":
			val, err := decoder.Token()
			if err != nil {
				return err
			}
			if val != nil {
				jtv, ok := val.(string)
				if !ok {
					return fmt.Errorf("expected AccessKeyType to be of type string, got %T instead", val)
				}
				sv.AccessKeyId = &jtv
			}

		case "expiration":
			val, err := decoder.Token()
			if err != nil {
				return err
			}
			if val != nil {
				jtv, ok := val.(json.Number)
				if !ok {
					return fmt.Errorf("expected ExpirationTimestampType to be json.Number, got %T instead", val)
				}
				i64, err := jtv.Int64()
				if err != nil {
					return err
				}
				sv.Expiration = &i64
			}

		case "secretAccessKey":
			val, err := decoder.Token()
			if err != nil {
				return err
			}
			if val != nil {
				jtv, ok := val.(string)
				if !ok {
					return fmt.Errorf("expected SecretAccessKeyType to be of type string, got %T instead", val)
				}
				sv.SecretAccessKey = &jtv
			}

		case "sessionToken":
			val, err := decoder.Token()
			if err != nil {
				return err
			}
			if val != nil {
				jtv, ok := val.(string)
				if !ok {
					return fmt.Errorf("expected SessionTokenType to be of type string, got %T instead", val)
				}
				sv.SessionToken = &jtv
			}

		default:
			err := smithyjson.DiscardUnknownField(decoder)
			if err != nil {
				return err
			}

		}
	}
	endToken, err := decoder.Token()
	if err != nil {
		return err
	}
	if t, ok := endToken.(json.Delim); !ok || t != '}' {
		return fmt.Errorf("expect `}` as end token")
	}

	*v = sv
	return nil
}

func awsRestjson1_deserializeDocumentRoleInfo(v **types.RoleInfo, decoder *json.Decoder) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	startToken, err := decoder.Token()
	if err == io.EOF {
		return nil
	}
	if err != nil {
		return err
	}
	if startToken == nil {
		return nil
	}
	if t, ok := startToken.(json.Delim); !ok || t != '{' {
		return fmt.Errorf("expect `{` as start token")
	}

	var sv *types.RoleInfo
	if *v == nil {
		sv = &types.RoleInfo{}
	} else {
		sv = *v
	}

	for decoder.More() {
		t, err := decoder.Token()
		if err != nil {
			return err
		}
		switch t {
		case "accountId":
			val, err := decoder.Token()
			if err != nil {
				return err
			}
			if val != nil {
				jtv, ok := val.(string)
				if !ok {
					return fmt.Errorf("expected AccountIdType to be of type string, got %T instead", val)
				}
				sv.AccountId = &jtv
			}

		case "roleName":
			val, err := decoder.Token()
			if err != nil {
				return err
			}
			if val != nil {
				jtv, ok := val.(string)
				if !ok {
					return fmt.Errorf("expected RoleNameType to be of type string, got %T instead", val)
				}
				sv.RoleName = &jtv
			}

		default:
			err := smithyjson.DiscardUnknownField(decoder)
			if err != nil {
				return err
			}

		}
	}
	endToken, err := decoder.Token()
	if err != nil {
		return err
	}
	if t, ok := endToken.(json.Delim); !ok || t != '}' {
		return fmt.Errorf("expect `}` as end token")
	}

	*v = sv
	return nil
}

func awsRestjson1_deserializeDocumentRoleListType(v *[]*types.RoleInfo, decoder *json.Decoder) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	startToken, err := decoder.Token()
	if err == io.EOF {
		return nil
	}
	if err != nil {
		return err
	}
	if startToken == nil {
		return nil
	}
	if t, ok := startToken.(json.Delim); !ok || t != '[' {
		return fmt.Errorf("expect `[` as start token")
	}

	var cv []*types.RoleInfo
	if *v == nil {
		cv = []*types.RoleInfo{}
	} else {
		cv = *v
	}

	for decoder.More() {
		var col *types.RoleInfo
		if err := awsRestjson1_deserializeDocumentRoleInfo(&col, decoder); err != nil {
			return err
		}
		cv = append(cv, col)

	}
	endToken, err := decoder.Token()
	if err != nil {
		return err
	}
	if t, ok := endToken.(json.Delim); !ok || t != ']' {
		return fmt.Errorf("expect `]` as end token")
	}

	*v = cv
	return nil
}

func awsRestjson1_deserializeDocumentTooManyRequestsException(v **types.TooManyRequestsException, decoder *json.Decoder) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	startToken, err := decoder.Token()
	if err == io.EOF {
		return nil
	}
	if err != nil {
		return err
	}
	if startToken == nil {
		return nil
	}
	if t, ok := startToken.(json.Delim); !ok || t != '{' {
		return fmt.Errorf("expect `{` as start token")
	}

	var sv *types.TooManyRequestsException
	if *v == nil {
		sv = &types.TooManyRequestsException{}
	} else {
		sv = *v
	}

	for decoder.More() {
		t, err := decoder.Token()
		if err != nil {
			return err
		}
		switch t {
		case "message":
			val, err := decoder.Token()
			if err != nil {
				return err
			}
			if val != nil {
				jtv, ok := val.(string)
				if !ok {
					return fmt.Errorf("expected ErrorDescription to be of type string, got %T instead", val)
				}
				sv.Message = &jtv
			}

		default:
			err := smithyjson.DiscardUnknownField(decoder)
			if err != nil {
				return err
			}

		}
	}
	endToken, err := decoder.Token()
	if err != nil {
		return err
	}
	if t, ok := endToken.(json.Delim); !ok || t != '}' {
		return fmt.Errorf("expect `}` as end token")
	}

	*v = sv
	return nil
}

func awsRestjson1_deserializeDocumentUnauthorizedException(v **types.UnauthorizedException, decoder *json.Decoder) error {
	if v == nil {
		return fmt.Errorf("unexpected nil of type %T", v)
	}
	startToken, err := decoder.Token()
	if err == io.EOF {
		return nil
	}
	if err != nil {
		return err
	}
	if startToken == nil {
		return nil
	}
	if t, ok := startToken.(json.Delim); !ok || t != '{' {
		return fmt.Errorf("expect `{` as start token")
	}

	var sv *types.UnauthorizedException
	if *v == nil {
		sv = &types.UnauthorizedException{}
	} else {
		sv = *v
	}

	for decoder.More() {
		t, err := decoder.Token()
		if err != nil {
			return err
		}
		switch t {
		case "message":
			val, err := decoder.Token()
			if err != nil {
				return err
			}
			if val != nil {
				jtv, ok := val.(string)
				if !ok {
					return fmt.Errorf("expected ErrorDescription to be of type string, got %T instead", val)
				}
				sv.Message = &jtv
			}

		default:
			err := smithyjson.DiscardUnknownField(decoder)
			if err != nil {
				return err
			}

		}
	}
	endToken, err := decoder.Token()
	if err != nil {
		return err
	}
	if t, ok := endToken.(json.Delim); !ok || t != '}' {
		return fmt.Errorf("expect `}` as end token")
	}

	*v = sv
	return nil
}