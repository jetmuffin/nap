package utils

import (
	"encoding/json"
	"net/http"
	"strings"
)

type validationError struct {
	cause error
}

func (e validationError) Error() string {
	return e.cause.Error()
}

func (e validationError) Cause() error {
	return e.cause
}

// WriteJSON writes the value v to the http response stream as json with standard json encoding.
func WriteJSON(w http.ResponseWriter, code int, v interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(false)
	return enc.Encode(v)
}

// WriteError writes the error message to the http response stream.
func WriteError(w http.ResponseWriter, err error) {
	statusCode := http.StatusBadRequest
	response := &ErrorResponse{
		Message: err.Error(),
	}
	WriteJSON(w, statusCode, response)
}

// ParseForm ensures the request form is parsed even with invalid content types.
// If we don't do this, POST method without Content-type (even with empty body) will fail.
func ParseForm(r *http.Request) error {
	if r == nil {
		return nil
	}
	if err := r.ParseForm(); err != nil && !strings.HasPrefix(err.Error(), "mime:") {
		return validationError{err}
	}
	return nil
}
