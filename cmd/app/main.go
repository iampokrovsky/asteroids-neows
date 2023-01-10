package main

import (
	"github.com/pokrovsky-io/neows-asteroids/config"
	"github.com/pokrovsky-io/neows-asteroids/internal/app"
)

func main() {
	cfg, _ := config.NewConfig()

	app.Run(cfg)
}
