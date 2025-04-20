package main

import (
	"blockchain_votation_system/config"
	"blockchain_votation_system/constants"
	"blockchain_votation_system/controllers"
	"blockchain_votation_system/services"
	"context"
	"log"

	"github.com/go-playground/validator"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

var Validate = *validator.New()

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.InitValidator()
	config.InitDB()

	ctx, cancel := context.WithTimeout(context.Background(), constants.ContextTimeout)
	defer cancel()

	services.GenerateGenesisBlock(ctx)

	e := echo.New()

	controllers.RegisterHealthRoutes(e)

	apiGroup := e.Group("/api")
	controllers.RegisterVotersRoutes(apiGroup)
	controllers.RegisterBlockchainRoutes(apiGroup)

	e.Logger.Fatal(e.Start(":1323"))
}
