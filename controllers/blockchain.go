package controllers

import (
	"blockchain_votation_system/constants"
	"blockchain_votation_system/entities"
	"blockchain_votation_system/services"
	"blockchain_votation_system/utils"
	"context"
	"strconv"

	"github.com/labstack/echo/v4"
)

func RegisterBlockchainRoutes(apiGroup *echo.Group) {
	/*
		GET ALL THE BLOCKS
	*/
	apiGroup.GET("/blockchain", func(c echo.Context) error {
		ctx, cancel := context.WithTimeout(c.Request().Context(), constants.ContextTimeout)
		defer cancel()

		msg, err := services.GetAllBlocks(ctx)

		if err != nil {
			return utils.ServerErrorResponse(c, echo.Map{"error": err.Error()})
		}

		return utils.SuccessResponse(c, msg)
	})

	/*
		ADD NEW BLOCK
	*/
	apiGroup.POST("/blockchain", func(c echo.Context) error {
		var block entities.Block

		if err := c.Bind(&block); err != nil {
			return utils.BadRequestErrorResponse(c, map[string]string{"error": err.Error()})
		}

		ctx, cancel := context.WithTimeout(c.Request().Context(), constants.ContextTimeout)
		defer cancel()

		msg, err := services.AddBlock(ctx, &block)

		if err != nil {
			return utils.ServerErrorResponse(c, echo.Map{"error": err.Error()})
		}

		return utils.SuccessResponse(c, msg)
	})

	/*
		VALIDATE A BLOCK
	*/
	apiGroup.GET("/blockchain/:index/validate", func(c echo.Context) error {
		indexParam := c.Param("index")
		index, err := strconv.Atoi(indexParam)

		if err != nil {
			return utils.BadRequestErrorResponse(c, echo.Map{"error": "invalid index"})
		}

		ctx, cancel := context.WithTimeout(c.Request().Context(), constants.ContextTimeout)
		defer cancel()

		msg, err := services.ValidateBlock(ctx, index)

		if err != nil {
			return utils.ServerErrorResponse(c, echo.Map{"error": err.Error()})
		}

		return utils.SuccessResponse(c, msg)
	})
}
