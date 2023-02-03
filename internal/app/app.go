package app

import (
	"asteroids-neows/config"
	"asteroids-neows/internal/repo/neows"
	"asteroids-neows/internal/repo/storage"
	"asteroids-neows/internal/transport"
	"asteroids-neows/internal/usecase"
	"asteroids-neows/pkg/httpserver"
	"asteroids-neows/pkg/postgres"
	"github.com/gin-gonic/gin"
	"log"
)

func Run(cfg *config.Config) {
	// DB
	db, err := postgres.New(cfg.DB.GetURL())
	if err != nil {
		log.Fatal(err)
	}

	// Storage
	stg, err := storage.New(db)
	if err != nil {
		log.Fatal(err)
	}

	// NeoWs API
	api := neows.New(cfg.NeoWs.URL, cfg.NeoWs.ApiKey)

	// Use case
	uc := usecase.New(stg, api)

	// HTTP Server
	handler := gin.New()
	transport.NewRouter(handler, uc)
	httpServer := httpserver.New(handler, cfg.HTTP.Port)

	if err = httpServer.Run(); err != nil {
		log.Fatal(err)
	}
}
