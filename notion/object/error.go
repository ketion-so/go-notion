package object

type ErrorCode string

const (
	InvalidJsonErrorCode        ErrorCode = "invalid_json"
	InvalidRequestURLErrorCode  ErrorCode = "invalid_request_url"
	InvalidRequestErrorCode     ErrorCode = "Invalid_request"
	ValidationErroreErrorCode   ErrorCode = "validation_error"
	UnauthorizedCode            ErrorCode = "invalid_json"
	RestrictedResourceErrorCode ErrorCode = "restricted_resource"
	ObjectNotFoundErrorCode     ErrorCode = "object_not_found"
	ConflictErrorErrorCode      ErrorCode = "invalid_json"
	RateLimitedErrorCode        ErrorCode = "conflict_error"
	InternalServerErrorCode     ErrorCode = "rate_limited"
	ServiceUnavailableErrorCode ErrorCode = "internal_server_error"
)
