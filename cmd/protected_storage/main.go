package main

import (
	"log"

	"protected-storage-server/internal/app"
	"protected-storage-server/internal/config"
)

func main() {
	var cfg config.Config
	err := cfg.Init()
	if err != nil {
		log.Fatal(err)
	}

	grpcApp, err := app.NewGRPC(cfg)
	if err != nil {
		log.Fatal(err)
	}

	err = grpcApp.Run(cfg)
	if err != nil {
		log.Fatal(err)
	}
}
