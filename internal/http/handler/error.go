package handler

import (
	"net/http"

	"github.com/pkg/errors"

	"github.com/derangga/shopifyx/internal/http/response"
	errorpkg "github.com/derangga/shopifyx/internal/pkg/error"
	"github.com/labstack/echo/v4"
)

func NewCustomErrorResponse(c echo.Context, err error) error {
	if err == nil {
		return nil
	}

	var custErr errorpkg.CustomError
	if !errors.As(err, &custErr) {
		custErr = errorpkg.CustomError{
			Message:  http.StatusText(http.StatusInternalServerError),
			HTTPCode: http.StatusInternalServerError,
			Err:      err,
		}
	}

	return c.JSON(custErr.HTTPCode, response.BaseResponse{
		Message: custErr.Message,
	})
}
