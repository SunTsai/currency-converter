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

	exchangeRates, err := util.LoadExchangeRates(".")
	if err != nil {
		log.Fatal("Failed to load the exchange rates: ", err)
	}

	server, err := api.NewServer(exchangeRates)
	if err != nil {
		log.Fatal("Failed to create the server: ", err)
	}

	if err := server.Start(config.ServerAddress); err != nil {
		log.Fatal("Failed to start the server: ", err)
	}
}
