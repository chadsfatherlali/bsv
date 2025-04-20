package main

import (
	"blockchain_votation_system/config"
	"blockchain_votation_system/constants"
	"blockchain_votation_system/controllers"
	"blockchain_votation_system/services"
	"context"

	"github.com/labstack/echo/v4"
)

func main() {
	config.LoadEnvs()
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
