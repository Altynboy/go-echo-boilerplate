package wrapper

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type SuccessResponse struct {
	Success bool `json:"Success"`
	Message string      `json:"Message"`
	Data    interface{} `json:"Data"`
}

type ErrorResponse struct {
	Success bool   `json:"Success"`
	Message string `json:"Message"`
}

func newSuccessResponse(msg string, data interface{}) SuccessResponse {
	return SuccessResponse{
		Success: true,
		Message: msg,
		Data:    data,
	}
}

func newErrorResponse(msg string) ErrorResponse {
	return ErrorResponse{
		Success: false,
		Message: msg,
	}
}

func Response(c echo.Context, status int, msg string, data ...interface{}) error {
	if status == http.StatusOK {
		return c.JSON(status, newSuccessResponse(msg, data))
	} else {
		return c.JSON(status, newErrorResponse(msg))
	}
}