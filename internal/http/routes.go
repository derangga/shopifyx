package http

import (
	"github.com/derangga/shopifyx/internal/http/handler"
	"github.com/labstack/echo/v4"
)

func RegisterRoute(
	e *echo.Echo,
	h *handler.Handlers,
	jwtMiddleware echo.MiddlewareFunc,
) {
	v1 := e.Group("/v1")
	v1.POST("/user/register", h.AuthHandler.Register)
	v1.POST("/user/login", h.AuthHandler.Login)

	v1Auth := e.Group("/v1", jwtMiddleware)
	v1Auth.POST("/bank/account", h.BankHandler.Create)
}
