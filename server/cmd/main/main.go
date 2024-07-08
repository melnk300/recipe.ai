package main

import (
	"github.com/melnk300/recipe.ai/server/internal/server"
	"github.com/melnk300/recipe.ai/server/pkg/config"
	"log"
)

func main() {
	config.Configure("config.yml")

	log.Println("Server started successful on port " + config.Config.Server.Port)

	if err := server.InitServer(config.Config.Server); err != nil {
		log.Fatal("init server failed")
	}
}
