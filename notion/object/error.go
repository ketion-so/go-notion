package object

// ErrorCode represents the Notion API errors.
// See details: https://developers.notion.com/reference/errors
type ErrorCode string

const (
	ErrInvalidJson        ErrorCode = "invalid_json"
	ErrInvalidRequestURL  ErrorCode = "invalid_request_url"
	ErrInvalidRequest     ErrorCode = "Invalid_request"
	ErrValidationErrore   ErrorCode = "validation_error"
	ErrUnautho            ErrorCode = "invalid_json"
	ErrRestrictedResource ErrorCode = "restricted_resource"
	ErrObjectNotFound     ErrorCode = "object_not_found"
	ErrConflictError      ErrorCode = "invalid_json"
	ErrRateLimited        ErrorCode = "conflict_error"
	ErrInternalServer     ErrorCode = "rate_limited"
	ErrServiceUnavailable ErrorCode = "internal_server_error"
)
