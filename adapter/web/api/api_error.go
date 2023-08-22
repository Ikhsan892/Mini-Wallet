package api

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strings"
)

type Data struct {
	Message string `json:"error"`
}

// ApiError defines the struct for a basic api error response.
type ApiError struct {
	Code   int    `json:"-"`
	Status string `json:"status"`
	Data   Data   `json:"data"`

	// stores unformatted error data (could be an internal error, text, etc.)
	rawData any
}

// Error makes it compatible with the `error` interface.
func (e *ApiError) Error() string {
	return e.Data.Message
}

// RawData returns the unformatted error data (could be an internal error, text, etc.)
func (e *ApiError) RawData() any {
	return e.rawData
}

// NewNotFoundError creates and returns 404 `ApiError`.
func NewNotFoundError(message string, data any) *ApiError {
	if message == "" {
		message = "The requested resource wasn't found."
	}

	return NewApiError(http.StatusNotFound, message, data)
}

// NewBadRequestError creates and returns 400 `ApiError`.
func NewBadRequestError(message string, data any) *ApiError {
	if message == "" {
		message = "Something went wrong while processing your request."
	}

	return NewApiError(http.StatusBadRequest, message, data)
}

// NewForbiddenError creates and returns 403 `ApiError`.
func NewForbiddenError(message string, data any) *ApiError {
	if message == "" {
		message = "You are not allowed to perform this request."
	}

	return NewApiError(http.StatusForbidden, message, data)
}

// NewUnauthorizedError creates and returns 401 `ApiError`.
func NewUnauthorizedError(message string, data any) *ApiError {
	if message == "" {
		message = "Missing or invalid authentication token."
	}

	return NewApiError(http.StatusUnauthorized, message, data)
}

// NewApiError creates and returns new normalized `ApiError` instance.
func NewApiError(status int, message string, data any) *ApiError {
	return &ApiError{
		Code:   status,
		Status: "fail",
		Data:   Data{strings.TrimSpace(message)},
	}
}

func NewValidationError(err error) *ApiError {
	message := ""

	if castedObject, ok := err.(validator.ValidationErrors); ok {
		for _, err := range castedObject {
			switch err.Tag() {
			case "required":
				message = fmt.Sprintf("%s is required",
					err.Field())
			case "email":
				message = fmt.Sprintf("%s is not valid email",
					err.Field())
			case "gte":
				message = fmt.Sprintf("%s value must be greater than %s",
					err.Field(), err.Param())
			case "lte":
				message = fmt.Sprintf("%s value must be lower than %s",
					err.Field(), err.Param())
			case "min":
				message = fmt.Sprintf("%s value at least %s data", err.Field(), err.Param())
			}
			break
		}
	}

	return NewApiError(http.StatusUnprocessableEntity, message, err)
}
