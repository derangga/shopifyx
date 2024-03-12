// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package app

import (
	"github.com/derangga/shopifyx/internal/config"
	"github.com/derangga/shopifyx/internal/http"
	"github.com/derangga/shopifyx/internal/http/handler"
	"github.com/derangga/shopifyx/internal/repository"
	"github.com/derangga/shopifyx/internal/usecase"
)

// Injectors from wire.go:

func InitHTTPServer(cfg *config.Config) http.HttpServer {
	databaseConfig := provideDBConfig(cfg)
	db := provideDB(databaseConfig)
	userRepository := repository.NewUserRepository(db)
	unitOfWork := repository.NewUnitOfWork(db)
	authConfig := provideAuthConfig(cfg)
	authUsecase := usecase.NewAuthUsecase(userRepository, unitOfWork, authConfig)
	validate := provideValidator()
	authHandler := handler.NewAuthHandler(authUsecase, validate)
	handlers := handler.NewHandlers(authHandler)
	middlewareFunc := provideJWTMiddleware(authConfig)
	httpServerItf := http.NewHttpServer(cfg, handlers, middlewareFunc)
	return httpServerItf
}