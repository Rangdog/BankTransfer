package errorcode

const (
	Success = 0

	// Validation errors (1000-1999)
	ErrMissingParameter = 1001
	ErrInvalidParameter = 1002

	// Auth errors (2000-2999)
	ErrUnauthorized = 2001
	ErrForbidden    = 2002

	// Business logic errors (3000-3999)
	ErrUserNotFound    = 3001
	ErrOrderNotAllowed = 3002

	// External service errors (4000-4999)
	ErrExternalAPI = 4001
	ErrDatabase    = 4002

	// System errors (5000-5999)
	ErrInternalServer = 5000
)
