package controllers

import (
	"BiometricToken/db"
	"BiometricToken/lib"
	Userlib "BiometricToken/lib/user"
	"BiometricToken/models"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

func RegisterUsers(c echo.Context)	error  {
	db := db.DbManager()
	params := new(Userlib.RegisterParams)

	if err := c.Bind(params); err !=nil {
		return BadRequestResponse(c,lib.INVALID_BODY)
	}
	user := new(models.User)
	exists := db.Where("email = ?", params.Email).Find(&user).RecordNotFound()
	if !exists {
		return BadRequestResponse(c,lib.AccountExists)
	}
	user.Email	=	params.Email
	user.FullName	=	params.FullName
	user.CreatedAt	=	time.Now()
	user.BioAuth	=	params.BioAuth
	user.Active	=	true
	user.Password	=	lib.GenerateHashFromPassword(params.Password)
	db.Save(&user)
	return DataResponse(c, user, http.StatusOK)
}