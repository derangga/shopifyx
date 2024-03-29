package http

import (
	"context"
	"log"

	"github.com/derangga/shopifyx/internal/config"
	"github.com/derangga/shopifyx/internal/http/handler"
	"github.com/derangga/shopifyx/internal/http/middleware"
	"github.com/labstack/echo/v4"
	"github.com/prometheus/client_golang/prometheus"
)

type HttpServer interface {
	ListenAndServe() error
	Stop()
}

type Server struct {
	echo         *echo.Echo
	cfg          *config.Config
	handler      *handler.Handlers
	jwtAuth      *middleware.JWTAuth
	reqHistogram *prometheus.HistogramVec
}

func NewHttpServer(
	cfg *config.Config,
	handler *handler.Handlers,
	jwtAuth *middleware.JWTAuth,
	rh *prometheus.HistogramVec,
) HttpServer {
	e := echo.New()
	middleware.SetupGlobalMiddleware(e, cfg.Application)

	srv := &Server{
		echo:         e,
		cfg:          cfg,
		handler:      handler,
		jwtAuth:      jwtAuth,
		reqHistogram: rh,
	}

	srv.connectCoreWithEcho()
	return srv
}

func (s *Server) ListenAndServe() error {
	return s.echo.Start(":" + s.cfg.Application.Port)
}

func (s *Server) Stop() {
	e := s.echo
	err := e.Server.Shutdown(context.Background())
	if err != nil {
		log.Fatal("failed to open shutdown service:", err.Error())
	}
}

func (s *Server) connectCoreWithEcho() {
	RegisterRoute(s.echo, s.handler, s.jwtAuth, s.reqHistogram)
}
