package http

import (
	"net/http"
	"strconv"
	"time"

	"github.com/derangga/shopifyx/internal/http/handler"
	"github.com/derangga/shopifyx/internal/http/middleware"
	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var reqHistogram *prometheus.HistogramVec

func RegisterRoute(
	e *echo.Echo,
	h *handler.Handlers,
	jwtAuth *middleware.JWTAuth,
	rh *prometheus.HistogramVec,
) {
	authMiddleware := jwtAuth.ToMiddleware()
	reqHistogram = rh

	// prometheus metrics
	e.GET("/metrics", echo.WrapHandler(promhttp.Handler()))

	// v1 route grouping
	v1 := e.Group("/v1")

	// user management
	newRoute(v1, http.MethodPost, "/user/register", h.AuthHandler.Register)
	newRoute(v1, http.MethodPost, "/user/login", h.AuthHandler.Login)

	// bank account
	newRoute(v1, http.MethodGet, "/bank/account", h.BankHandler.Fetch, authMiddleware)
	newRoute(v1, http.MethodPost, "/bank/account", h.BankHandler.Create, authMiddleware)
	newRoute(v1, http.MethodPatch, "/bank/account/:bankAccountId", h.BankHandler.Update, authMiddleware)
	newRoute(v1, http.MethodDelete, "/bank/account/:bankAccountId", h.BankHandler.Delete, authMiddleware)

	// product
	newRoute(v1, http.MethodGet, "/product", h.ProductHandler.Fetch, authMiddleware)
	newRoute(v1, http.MethodPost, "/product", h.ProductHandler.Create, authMiddleware)
	newRoute(v1, http.MethodPatch, "/product/:id", h.ProductHandler.Update, authMiddleware)
	newRoute(v1, http.MethodDelete, "/product/:id", h.ProductHandler.Delete, authMiddleware)
	newRoute(v1, http.MethodPost, "/product/:id/stock", h.ProductHandler.UpdateStock, authMiddleware)

	// image upload
	newRoute(v1, http.MethodPost, "/image", h.ImageHandler.Upload, authMiddleware)

	// payment
	newRoute(v1, http.MethodPost, "/product/:id/buy", h.PaymentHandler.Create, authMiddleware)
}

func newRoute(e *echo.Group, method string, path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) {
	e.Add(method, path, wrapHandlerWithMetrics(path, method, h), m...)
}

func wrapHandlerWithMetrics(path string, method string, handler echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		startTime := time.Now()

		// Execute the actual handler and catch any errors
		err := handler(c)

		// Regardless of whether an error occurred, record the metrics
		duration := time.Since(startTime).Seconds()

		reqHistogram.WithLabelValues(path, method, strconv.Itoa(c.Response().Status)).Observe(duration)

		return err
	}
}
