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

	// user management
	v1.POST("/user/register", h.AuthHandler.Register)
	v1.POST("/user/login", h.AuthHandler.Login)

	// bank account
	v1.POST("/bank/account", h.BankHandler.Create, authMiddleware)
	v1.PATCH("/bank/account/:bankAccountId", h.BankHandler.Update, authMiddleware)

	// product management
	v1.POST("/product", h.ProductHandler.Create, authMiddleware)
}
