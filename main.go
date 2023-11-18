package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"suntsai/currency-converter/api"
	"suntsai/currency-converter/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Failed to load the config: ", err)
	}
	gin.SetMode(config.GinMode)

	rates := api.NewExchangeRates()
	server, err := api.NewServer(rates)
	if err != nil {
		log.Fatal("Failed to create the server: ", err)
	}

	if err := server.Start(config.ServerAddress); err != nil {
		log.Fatal("Failed to start the server: ", err)
	}
}
