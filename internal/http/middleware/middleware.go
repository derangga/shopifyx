package middleware

import (
	"github.com/derangga/shopifyx/internal/config"
	"github.com/labstack/echo/v4"
	emidd "github.com/labstack/echo/v4/middleware"
)

func SetupGlobalMiddleware(e *echo.Echo, cfg config.ApplicationConfig) {
	e.Use(emidd.ContextTimeoutWithConfig(emidd.ContextTimeoutConfig{Skipper: emidd.DefaultSkipper, Timeout: cfg.Timeout}))
	e.Use(emidd.Logger())
	e.Use(emidd.Recover())
}
