package router

import (
	"BiometricToken/controllers"
	"github.com/labstack/echo/v4"
)

func home(c echo.Context) error {

	return c.JSON(200, "hoome")
}

func New() *echo.Echo {

	e := echo.New()
	e.GET("/index", home)
	e.GET("/getusers", controllers.GetUsers)
	e.POST("/registeruser", controllers.RegisterUsers)

	return e
}
