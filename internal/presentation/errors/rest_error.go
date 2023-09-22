package errors

import "net/http"

type RestError struct {
	Message string  `json:"message"`
	Err     string  `json:"error"`
	Code    int     `json:"code"`
	Causes  []Cause `json:"causes,omitempty"`
}

type Cause struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func NewRestErr(message string, err string, code int, causes []Cause) *RestError {
	return &RestError{
		Message: message,
		Err:     err,
		Code:    code,
		Causes:  causes,
	}
}

func NewBadRequestError(message string) *RestError {
	return &RestError{
		Message: message,
		Err:     "bad_request",
		Code:    http.StatusBadRequest,
	}
}

func NewBadRequestValidationError(message string, causes []Cause) *RestError {
	return &RestError{
		Message: message,
		Err:     "bad_request",
		Code:    http.StatusBadRequest,
		Causes:  causes,
	}
}

func NewInternalServerError(message string) *RestError {
	return &RestError{
		Message: message,
		Err:     "internal_server_error",
		Code:    http.StatusInternalServerError,
	}
}

func NewForbiddenError(message string) *RestError {
	return &RestError{
		Message: message,
		Err:     "forbidden",
		Code:    http.StatusForbidden,
	}
}

func NewNotFoundError(message string) *RestError {
	return &RestError{
		Message: message,
		Err:     "not_found",
		Code:    http.StatusNotFound,
	}
}

func NewUnauthorizedError(message string) *RestError {
	return &RestError{
		Message: message,
		Err:     "unauthorized",
		Code:    http.StatusUnauthorized,
	}
}

func (r *RestError) GetMessage() string {
	return r.Message
}
