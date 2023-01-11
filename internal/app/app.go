package app

import (
	"github.com/gin-gonic/gin"
	"github.com/pokrovsky-io/neows-asteroids/config"
	"github.com/pokrovsky-io/neows-asteroids/internal/repo/neows"
	"github.com/pokrovsky-io/neows-asteroids/internal/repo/storage"
	"github.com/pokrovsky-io/neows-asteroids/internal/transport"
	"github.com/pokrovsky-io/neows-asteroids/internal/usecase"
	"github.com/pokrovsky-io/neows-asteroids/pkg/httpserver"
	"github.com/pokrovsky-io/neows-asteroids/pkg/postgres"
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
