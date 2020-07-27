package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

// Error Struct
type Error struct {
	Success bool        `json:"success"`
	Message interface{} `json:"message"`
}

type Data struct {
	Success	bool	`json:"success"`
	Data	interface{}	`json:"data"`
}

func BadRequestResponse(c echo.Context, message string) error {
	return c.JSONPretty(http.StatusBadRequest, Error{
		Message: message,
		Success: false,
	}, " ")
}

func DataResponse(c echo.Context, data interface{}, status int) error {
	return c.JSONPretty(status, Data{
		Success: true,
		Data:    data,
	},
	" ")
}

