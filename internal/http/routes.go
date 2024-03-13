package http

import (
	"github.com/derangga/shopifyx/internal/http/handler"
	"github.com/derangga/shopifyx/internal/http/middleware"
	"github.com/labstack/echo/v4"
)

func RegisterRoute(
	e *echo.Echo,
	h *handler.Handlers,
	jwtAuth *middleware.JWTAuth,
) {
	authMiddleware := jwtAuth.ToMiddleware()

	v1 := e.Group("/v1")
	v1.POST("/user/register", h.AuthHandler.Register)
	v1.POST("/user/login", h.AuthHandler.Login)

	v1.POST("/product", h.ProductHandler.Create, authMiddleware)
}
