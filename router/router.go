package router

import (
	"BiometricToken/db"
	"github.com/labstack/echo/v4"
)

func home(c echo.Context) error {

	return c.JSON(200, "hoome")
}

func New() *echo.Echo {

	e := echo.New()
	e.GET("/index", home)
	db.DbInit()
	return e
}
