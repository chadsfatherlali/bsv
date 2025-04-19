package controllers

import (
	"blockchain_votation_system/constants"
	"blockchain_votation_system/entities"
	"blockchain_votation_system/services"
	"context"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func RegisterBlockchainRoutes(e *echo.Echo) {
	blockchainRoutes := e.Group("/api")

	/*
		GET ALL THE BLOCKS
	*/
	blockchainRoutes.GET("/blockchain", func(c echo.Context) error {
		ctx, cancel := context.WithTimeout(c.Request().Context(), constants.ContextTimeout)
		defer cancel()

		msg, err := services.GetAllBlocks(ctx)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
		}

		return c.JSON(http.StatusOK, msg)
	})

	/*
		ADD NEW BLOCK
	*/
	blockchainRoutes.POST("/blockchain", func(c echo.Context) error {
		var block entities.Block

		if err := c.Bind(&block); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		}

		ctx, cancel := context.WithTimeout(c.Request().Context(), constants.ContextTimeout)
		defer cancel()

		msg, err := services.AddBlock(ctx, &block)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
		}

		return c.JSON(http.StatusOK, msg)
	})

	/*
		VALIDATE A BLOCK
	*/
	blockchainRoutes.GET("/blockchain/:index/validate", func(c echo.Context) error {
		indexParam := c.Param("index")
		index, err := strconv.Atoi(indexParam)

		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid index"})
		}

		ctx, cancel := context.WithTimeout(c.Request().Context(), constants.ContextTimeout)
		defer cancel()

		msg, err := services.ValidateBlock(ctx, index)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
		}

		return c.JSON(http.StatusOK, msg)
	})
}
