package app

import (
	"github.com/derangga/shopifyx/internal/config"
	"github.com/derangga/shopifyx/internal/http"
	"github.com/derangga/shopifyx/internal/http/handler"
	"github.com/derangga/shopifyx/internal/http/middleware"
	"github.com/derangga/shopifyx/internal/pkg/database"
	"github.com/derangga/shopifyx/internal/repository"
	"github.com/derangga/shopifyx/internal/usecase"
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

var (
	repositoriesSet = wire.NewSet(
		repository.NewUserRepository,
		repository.NewUnitOfWork,
		repository.NewBankRepository,
	)

	usecasesSet = wire.NewSet(
		usecase.NewAuthUsecase,
		usecase.NewBankUsecase,
	)

	handlerSet = wire.NewSet(
		handler.NewAuthHandler,
		handler.NewBankHandler,
	)

	middlewareSet = wire.NewSet(
		provideJWTMiddleware,
	)

	allSet = wire.NewSet(
		provideAuthConfig,
		provideDBConfig,
		provideDB,
		provideValidator,
		repositoriesSet,
		usecasesSet,
		handlerSet,
		middlewareSet,
		handler.NewHandlers,
		http.NewHttpServer,
	)
)

func provideAuthConfig(cfg *config.Config) *config.AuthConfig {
	return &cfg.Auth
}

func provideDBConfig(cfg *config.Config) *config.DatabaseConfig {
	return &cfg.Database
}

func provideDB(cfg *config.DatabaseConfig) *sqlx.DB {
	return database.NewPostgresDatabase(cfg)
}

func provideValidator() *validator.Validate {
	return validator.New()
}

func provideJWTMiddleware(cfg *config.AuthConfig) echo.MiddlewareFunc {
	return middleware.NewJWTAuthMiddleware(cfg.JWTSecret)
}
