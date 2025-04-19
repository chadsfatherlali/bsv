package utils

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

/*
CUSTOM ERRORS
*/
type ApiResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

func SuccessResponse(c echo.Context, data any) error {
	return c.JSON(http.StatusOK, ApiResponse{
		Success: true,
		Data:    data,
	})
}

func BadRequestErrorResponse(c echo.Context, data any) error {
	return c.JSON(http.StatusBadRequest, ApiResponse{
		Success: false,
		Data:    data,
	})
}

func ServerErrorResponse(c echo.Context, data any) error {
	return c.JSON(http.StatusInternalServerError, ApiResponse{
		Success: false,
		Data:    data,
	})
}

/*
MESSAGES ERROR
*/
type VoterAlreadyExistsError struct {
	Document int
}

func (e *VoterAlreadyExistsError) Error() string {
	return fmt.Sprintf("Voter with document '%d' already exists", e.Document)
}
