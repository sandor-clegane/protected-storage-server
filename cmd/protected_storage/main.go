package main

import (
	"fmt"
	"log"

	"protected-storage-server/internal/app"
	"protected-storage-server/internal/config"
)

func main() {
	var cfg config.Config
	cfg.Init()
	fmt.Println(cfg)

	grpcApp, err := app.NewGRPC(cfg)
	if err != nil {
		log.Fatal(err)
	}

	err = grpcApp.Run(cfg)
	if err != nil {
		log.Fatal(err)
	}
}
