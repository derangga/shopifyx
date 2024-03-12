package main

import (
	"log"

	"github.com/derangga/shopifyx/internal/app"
	"github.com/derangga/shopifyx/internal/config"
)

func main() {
	cfg := config.MustGet()
	httpServer := app.InitHTTPServer(cfg)

	err := httpServer.ListenAndServe()
	if err != nil {
		log.Fatal("failed to serve http:", err.Error())
	}
}
