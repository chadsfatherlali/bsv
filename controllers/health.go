package controllers

import (
	"blockchain_votation_system/services"

	"github.com/labstack/echo/v4"
)

func RegisterHealthRoutes(e *echo.Echo) {
	healthRoutes := e.Group("/health")

	healthRoutes.GET("", func(c echo.Context) error {
		code, msg := services.ResponseOK()

		return c.String(code, msg)
	})
}
