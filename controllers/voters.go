package controllers

import (
	"blockchain_votation_system/constants"
	"blockchain_votation_system/entities"
	"blockchain_votation_system/services"
	"blockchain_votation_system/utils"
	"context"

	"github.com/labstack/echo/v4"
)

func RegisterVotersRoutes(apiGroup *echo.Group) {
	/*
		ADD NEW USER
	*/
	apiGroup.POST("/voters", func(c echo.Context) error {
		var voter entities.Voter

		if err := c.Bind(&voter); err != nil {
			return utils.BadRequestErrorResponse(c, map[string]string{"error": err.Error()})
		}

		ctx, cancel := context.WithTimeout(c.Request().Context(), constants.ContextTimeout)
		defer cancel()

		msg, err := services.AddVoter(ctx, &voter)

		if err != nil {
			return utils.ServerErrorResponse(c, echo.Map{"error": err.Error()})
		}

		return utils.SuccessResponse(c, msg)
	})
}
