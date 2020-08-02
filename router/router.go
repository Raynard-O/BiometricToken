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
	api := e.Group("/v1")
	e.GET("/index", home)
	e.POST("/verify", controllers.Verify)

	e.GET("/getusers", controllers.GetUsers)

	// admin access
	adminGroup := api.Group("/admin")
	adminGroup.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: Hmac,
		SigningMethod: "HS512",
		Skipper: func(c echo.Context) bool {
			if c.Path() == "/v1/admin/adminlogin"	|| c.Path() == "/v1/admin/createadmin"	|| c.Path() == "/v1/admin/getadmins" {
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
