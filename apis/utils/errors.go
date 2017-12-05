package utils

// ErrorResponse is the wrapper for http response when some errors happened.
type ErrorResponse struct {

	// The error message.
	// Required: true
	Message string `json:"message"`
}
