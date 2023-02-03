package main

import (
	"asteroids-neows/config"
	"asteroids-neows/internal/app"
)

func main() {
	cfg, _ := config.NewConfig()

	app.Run(cfg)
}
