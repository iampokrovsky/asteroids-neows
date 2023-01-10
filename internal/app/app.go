package app

import (
	"github.com/gin-gonic/gin"
	"github.com/pokrovsky-io/neows-asteroids/config"
	"github.com/pokrovsky-io/neows-asteroids/internal/transport"
	"github.com/pokrovsky-io/neows-asteroids/pkg/httpserver"
)

func Run(cfg *config.Config) {

	// HTTP Server
	handler := gin.New()
	transport.NewRouter(handler, nil)
	httpServer := httpserver.New(handler, cfg.HTTP.Port)

	// TODO: Обработать ошибку
	httpServer.Run()
}
