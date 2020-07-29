package router

import (
	"BiometricToken/configuration"
	"BiometricToken/controllers"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func home(c echo.Context) error {

	return c.JSON(200, "hoome")
}

var Hmac = configuration.GetHmacSigningKey()
func New() *echo.Echo {

	e := echo.New()
	e.GET("/index", home)
	e.POST("/verify", controllers.Verify)

	e.GET("/getusers", controllers.GetUsers)

	// admin access
	adminGroup := e.Group("/admin")
	adminGroup.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: Hmac,
		SigningMethod: "HS512",
		Skipper: func(c echo.Context) bool {
			if c.Path() == "/admin/adminlogin"	|| c.Path() == "/admin/createadmin"	|| c.Path() == "/admin/getadmins" {
				return true
			}
			return false
		},
	}))
	adminGroup.GET("/getadmins", controllers.GetAdmin)
	adminGroup.POST("/createadmin", controllers.CreateAdmin)
	adminGroup.POST("/adminlogin", controllers.LoginAdmin)
	adminGroup.POST("/registeruser", controllers.RegisterUsers)

	return e
}
