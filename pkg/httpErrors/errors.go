package httpErrors

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

// Rest error interface
type RestErr interface {
	Status() int
	Error() string
	Causes() interface{}
}

// Rest error struct
type RestError struct {
	ErrStatus int         `json:"status,omitempty"`
	ErrError  string      `json:"error,omitempty"`
	ErrCause  interface{} `json:"cause,omitempty"`
}

// Error  Error() interface method
func (e RestError) Error() string {
	return fmt.Sprintf("status: %d - errors: %s - cause: %v", e.ErrStatus, e.ErrError, e.ErrCause)
}

// Error status
func (e RestError) Status() int {
	return e.ErrStatus
}

// RestError Causes
func (e RestError) Causes() interface{} {
	return e.ErrCause
}

// New Rest Error
func NewRestError(status int, err string, cause interface{}) RestErr {
	return RestError{
		ErrStatus: status,
		ErrError:  err,
		ErrCause:  cause,
	}
}

// New Bad Request Error
func NewBadRequestError(cause interface{}) RestErr {
	return RestError{
		ErrStatus: http.StatusBadRequest,
		ErrError:  ErrBadRequest,
		ErrCause:  cause,
	}
}

// New Not Found Error
func NewNotFoundError(cause interface{}) RestErr {
	return RestError{
		ErrStatus: http.StatusNotFound,
		ErrError:  ErrNotFound,
		ErrCause:  cause,
	}
}

// New Unauthorized Error
func NewUnauthorizedError(cause interface{}) RestErr {
	return RestError{
		ErrStatus: http.StatusUnauthorized,
		ErrError:  ErrUnauthorized,
		ErrCause:  cause,
	}
}

// New Forbidden Error
func NewForbiddenError(cause interface{}) RestErr {
	return RestError{
		ErrStatus: http.StatusForbidden,
		ErrError:  ErrForbidden,
		ErrCause:  cause,
	}
}

// New Internal Server Error
func NewInternalServerError(cause interface{}) RestErr {
	result := RestError{
		ErrStatus: http.StatusInternalServerError,
		ErrError:  ErrInternalServer,
		ErrCause:  cause,
	}
	return result
}

func ParseErrors(err error) RestErr {
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return NewRestError(http.StatusNotFound, ErrNotFound, err)
	case errors.Is(err, context.DeadlineExceeded):
		return NewRestError(http.StatusRequestTimeout, ErrRequestTimeout, err)
	case strings.Contains(err.Error(), "SQLSTATE"):
		return parseSqlErrors(err)
	case strings.Contains(err.Error(), "Field validation"):
		return parseValidatorError(err)
	case strings.Contains(err.Error(), "Unmarshal"):
		return NewRestError(http.StatusBadRequest, ErrBadRequest, err)
	case strings.Contains(err.Error(), "UUID"):
		return NewRestError(http.StatusBadRequest, err.Error(), err)
	case strings.Contains(strings.ToLower(err.Error()), "cookie"):
		return NewRestError(http.StatusUnauthorized, ErrUnauthorized, err)
	case strings.Contains(strings.ToLower(err.Error()), "token"):
		return NewRestError(http.StatusUnauthorized, ErrUnauthorized, err)
	case strings.Contains(strings.ToLower(err.Error()), "bcrypt"):
		return NewRestError(http.StatusBadRequest, ErrBadRequest, err)
	default:
		if restErr, ok := err.(RestErr); ok {
			return restErr
		}
		return NewInternalServerError(err)
	}
}

func parseSqlErrors(err error) RestErr {
	switch {
	case strings.Contains(err.Error(), "23505"):
		return NewRestError(http.StatusBadRequest, ErrEmailAlreadyExists, err)
	default:
		return NewRestError(http.StatusBadRequest, ErrBadRequest, err)
	}
}

func parseValidatorError(err error) RestErr {
	switch {
	case strings.Contains(err.Error(), "Password"):
		return NewRestError(http.StatusBadRequest, "Invalid password, min length 6", err)
	case strings.Contains(err.Error(), "Email"):
		return NewRestError(http.StatusBadRequest, "Invalid email", err)
	default:
		return NewRestError(http.StatusBadRequest, ErrBadRequest, err)
	}
}

// Error response
func ErrorResponse(err error) (int, interface{}) {
	return ParseErrors(err).Status(), ParseErrors(err)
}
