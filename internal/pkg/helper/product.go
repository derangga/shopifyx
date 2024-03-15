package helper

import (
	"fmt"
	"net/http"
	"strconv"

	errorpkg "github.com/derangga/shopifyx/internal/pkg/error"
	"github.com/labstack/echo/v4"
)

func GetIntQueryParams(c echo.Context, defValue int, key string) (int, error) {
	if c.QueryParam(key) == "" {
		return defValue, nil
	}

	val, err := strconv.Atoi(c.QueryParam(key))
	if err != nil {
		return 0, errorpkg.NewCustomMessageError(fmt.Sprintf("%s should be integer, got error: %v", key, err), http.StatusBadRequest, err)
	}

	if val > 0 {
		return val, nil
	}

	return defValue, nil
}
