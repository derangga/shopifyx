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
	v1.GET("/bank/account", h.BankHandler.Fetch, authMiddleware)
	v1.POST("/bank/account", h.BankHandler.Create, authMiddleware)
	v1.PATCH("/bank/account/:bankAccountId", h.BankHandler.Update, authMiddleware)
	v1.DELETE("/bank/account/:bankAccountId", h.BankHandler.Delete, authMiddleware)

	// product management
	v1.POST("/product", h.ProductHandler.Create, authMiddleware)
	v1.PATCH("/product/:id", h.ProductHandler.Update, authMiddleware)
	v1.DELETE("/product/:id", h.ProductHandler.Delete, authMiddleware)
	v1.POST("/product/:id/stock", h.ProductHandler.UpdateStock, authMiddleware)

	// image upload
	v1.POST("/image", h.ImageHandler.Upload, authMiddleware)
  // product page
	v1.GET("/product", h.ProductHandler.Fetch, authMiddleware)
  // payment
	v1.POST("/product/:id/buy", h.PaymentHandler.Create, authMiddleware)
}
