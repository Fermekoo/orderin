package main

import (
	"context"
	"log"

	"github.com/Fermekoo/go-kapster/api"
	"github.com/Fermekoo/go-kapster/utils"
)

func main() {
	config, err := utils.LoadConfig("./")
	if err != nil {
		log.Fatal("failed to setup config %w", err)
	}

	apiServer, err := api.NewApiServer(config)
	if err != nil {
		log.Fatal("cannot start http server")
	}
	apiServer.Start(config.HTTPServerPort, context.Background())
}
