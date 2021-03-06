// Code generated by smithy-go-codegen DO NOT EDIT.

package types

type AccessDeniedExceptionReason string

// Enum values for AccessDeniedExceptionReason
const (
	AccessDeniedExceptionReasonUnauthorized_account     AccessDeniedExceptionReason = "UNAUTHORIZED_ACCOUNT"
	AccessDeniedExceptionReasonDependency_access_denied AccessDeniedExceptionReason = "DEPENDENCY_ACCESS_DENIED"
)

// Values returns all known values for AccessDeniedExceptionReason. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (AccessDeniedExceptionReason) Values() []AccessDeniedExceptionReason {
	return []AccessDeniedExceptionReason{
		"UNAUTHORIZED_ACCOUNT",
		"DEPENDENCY_ACCESS_DENIED",
	}
}

type ChecksumAggregationMethod string

// Enum values for ChecksumAggregationMethod
const (
	ChecksumAggregationMethodChecksum_aggregation_linear ChecksumAggregationMethod = "LINEAR"
)

// Values returns all known values for ChecksumAggregationMethod. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (ChecksumAggregationMethod) Values() []ChecksumAggregationMethod {
	return []ChecksumAggregationMethod{
		"LINEAR",
	}
}

type ChecksumAlgorithm string

// Enum values for ChecksumAlgorithm
const (
	ChecksumAlgorithmChecksum_algorithm_sha256 ChecksumAlgorithm = "SHA256"
)

// Values returns all known values for ChecksumAlgorithm. Note that this can be
// expanded in the future, and so it is only as up to date as the client. The
// ordering of this slice is not guaranteed to be stable across updates.
func (ChecksumAlgorithm) Values() []ChecksumAlgorithm {
	return []ChecksumAlgorithm{
		"SHA256",
	}
}

type RequestThrottledExceptionReason string

// Enum values for RequestThrottledExceptionReason
const (
	RequestThrottledExceptionReasonAccount_throttled            RequestThrottledExceptionReason = "ACCOUNT_THROTTLED"
	RequestThrottledExceptionReasonDependency_request_throttled RequestThrottledExceptionReason = "DEPENDENCY_REQUEST_THROTTLED"
)

// Values returns all known values for RequestThrottledExceptionReason. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (RequestThrottledExceptionReason) Values() []RequestThrottledExceptionReason {
	return []RequestThrottledExceptionReason{
		"ACCOUNT_THROTTLED",
		"DEPENDENCY_REQUEST_THROTTLED",
	}
}

type ResourceNotFoundExceptionReason string

// Enum values for ResourceNotFoundExceptionReason
const (
	ResourceNotFoundExceptionReasonSnapshot_not_found            ResourceNotFoundExceptionReason = "SNAPSHOT_NOT_FOUND"
	ResourceNotFoundExceptionReasonDependency_resource_not_found ResourceNotFoundExceptionReason = "DEPENDENCY_RESOURCE_NOT_FOUND"
)

// Values returns all known values for ResourceNotFoundExceptionReason. Note that
// this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (ResourceNotFoundExceptionReason) Values() []ResourceNotFoundExceptionReason {
	return []ResourceNotFoundExceptionReason{
		"SNAPSHOT_NOT_FOUND",
		"DEPENDENCY_RESOURCE_NOT_FOUND",
	}
}

type ServiceQuotaExceededExceptionReason string

// Enum values for ServiceQuotaExceededExceptionReason
const (
	ServiceQuotaExceededExceptionReasonDependency_service_quota_exceeded ServiceQuotaExceededExceptionReason = "DEPENDENCY_SERVICE_QUOTA_EXCEEDED"
)

// Values returns all known values for ServiceQuotaExceededExceptionReason. Note
// that this can be expanded in the future, and so it is only as up to date as the
// client. The ordering of this slice is not guaranteed to be stable across
// updates.
func (ServiceQuotaExceededExceptionReason) Values() []ServiceQuotaExceededExceptionReason {
	return []ServiceQuotaExceededExceptionReason{
		"DEPENDENCY_SERVICE_QUOTA_EXCEEDED",
	}
}

type Status string

// Enum values for Status
const (
	StatusCompleted Status = "completed"
	StatusPending   Status = "pending"
	StatusError     Status = "error"
)

// Values returns all known values for Status. Note that this can be expanded in
// the future, and so it is only as up to date as the client. The ordering of this
// slice is not guaranteed to be stable across updates.
func (Status) Values() []Status {
	return []Status{
		"completed",
		"pending",
		"error",
	}
}

type ValidationExceptionReason string

// Enum values for ValidationExceptionReason
const (
	ValidationExceptionReasonInvalid_customer_key       ValidationExceptionReason = "INVALID_CUSTOMER_KEY"
	ValidationExceptionReasonInvalid_page_token         ValidationExceptionReason = "INVALID_PAGE_TOKEN"
	ValidationExceptionReasonInvalid_block_token        ValidationExceptionReason = "INVALID_BLOCK_TOKEN"
	ValidationExceptionReasonInvalid_snapshot_id        ValidationExceptionReason = "INVALID_SNAPSHOT_ID"
	ValidationExceptionReasonUnrelated_snapshots        ValidationExceptionReason = "UNRELATED_SNAPSHOTS"
	ValidationExceptionReasonInvalid_block              ValidationExceptionReason = "INVALID_BLOCK"
	ValidationExceptionReasonInvalid_content_encoding   ValidationExceptionReason = "INVALID_CONTENT_ENCODING"
	ValidationExceptionReasonInvalid_tag                ValidationExceptionReason = "INVALID_TAG"
	ValidationExceptionReasonInvalid_dependency_request ValidationExceptionReason = "INVALID_DEPENDENCY_REQUEST"
	ValidationExceptionReasonInvalid_parameter_value    ValidationExceptionReason = "INVALID_PARAMETER_VALUE"
	ValidationExceptionReasonInvalid_volume_size        ValidationExceptionReason = "INVALID_VOLUME_SIZE"
)

// Values returns all known values for ValidationExceptionReason. Note that this
// can be expanded in the future, and so it is only as up to date as the client.
// The ordering of this slice is not guaranteed to be stable across updates.
func (ValidationExceptionReason) Values() []ValidationExceptionReason {
	return []ValidationExceptionReason{
		"INVALID_CUSTOMER_KEY",
		"INVALID_PAGE_TOKEN",
		"INVALID_BLOCK_TOKEN",
		"INVALID_SNAPSHOT_ID",
		"UNRELATED_SNAPSHOTS",
		"INVALID_BLOCK",
		"INVALID_CONTENT_ENCODING",
		"INVALID_TAG",
		"INVALID_DEPENDENCY_REQUEST",
		"INVALID_PARAMETER_VALUE",
		"INVALID_VOLUME_SIZE",
	}
}
