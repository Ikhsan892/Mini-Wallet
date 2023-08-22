package api

import (
	"github.com/labstack/echo/v4"
)

type ApiResponse struct {
	Status string `json:"status"`
	Data   any    `json:"data"`
}

func NewApiResponse(code int, data any, c echo.Context) error {
	resp := ApiResponse{
		Status: "success",
		Data:   data,
	}

	return c.JSON(code, resp)
}
