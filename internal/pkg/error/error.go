package error

import (
	"net/http"

	"github.com/pkg/errors"
)

// CustomError holds data for customized error
type CustomError struct {
	Message  string
	HTTPCode int
	Err      error
}

func (c CustomError) Error() string {
	if c.Err != nil {
		return c.Message + ": " + c.Err.Error()
	}

	return c.Message
}

func NewCustomError(httpCode int, err error) error {
	newErr := CustomError{
		Message:  http.StatusText(httpCode),
		HTTPCode: httpCode,
		Err:      err,
	}

	return errors.WithStack(newErr)
}

func NewCustomMessageError(message string, httpCode int, err error) error {
	newErr := CustomError{
		Message:  message,
		HTTPCode: httpCode,
		Err:      err,
	}

	return errors.WithStack(newErr)
}

type RowNotFound struct {
	Message string
}

func (r RowNotFound) Error() string {
	return r.Message
}
