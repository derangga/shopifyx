// Wire something

//go:build wireinject

package app

import (
	"github.com/derangga/shopifyx/internal/config"
	"github.com/derangga/shopifyx/internal/http"
	"github.com/google/wire"
)

func InitHTTPServer(cfg *config.Config) http.HttpServerItf {
	wire.Build(allSet)
	return &http.HttpServer{}
}
