package main

import (
	"context"
	"log"

	"github.com/Fermekoo/orderin-api/api"
	"github.com/Fermekoo/orderin-api/utils"
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
