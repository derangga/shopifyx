package http

import (
	"context"
	"log"

	"github.com/derangga/shopifyx/internal/config"
	"github.com/derangga/shopifyx/internal/http/handler"
	"github.com/derangga/shopifyx/internal/http/middleware"
	"github.com/labstack/echo/v4"
)

type HttpServer interface {
	ListenAndServe() error
	Stop()
}

type httpServer struct {
	echo          *echo.Echo
	cfg           *config.Config
	handler       *handler.Handlers
	jwtMiddleware echo.MiddlewareFunc
}

// New HTTP Server
func NewHttpServer(
	cfg *config.Config,
	handler *handler.Handlers,
	jwtMiddleware echo.MiddlewareFunc,
) HttpServer {
	e := echo.New()
	middleware.SetupGlobalMiddleware(e, cfg.Application)

	srv := &httpServer{
		echo:          e,
		cfg:           cfg,
		handler:       handler,
		jwtMiddleware: jwtMiddleware,
	}

	srv.connectCoreWithEcho()
	return srv
}

func (h *httpServer) ListenAndServe() error {
	return h.echo.Start(":" + h.cfg.Application.Port)
}

func (h *httpServer) Stop() {
	e := h.echo
	err := e.Server.Shutdown(context.Background())
	if err != nil {
		log.Fatal("failed to open shutdown service:", err.Error())
	}
}

func (h *httpServer) connectCoreWithEcho() {
	RegisterRoute(h.echo, h.handler, h.jwtMiddleware)
}
