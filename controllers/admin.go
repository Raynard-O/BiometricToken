package controllers

import (
	"BiometricToken/db"
	"BiometricToken/lib"
	Adminlib "BiometricToken/lib/admin"
	"BiometricToken/models"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"time"
)

func AdminLogin(c echo.Context) error {
	db := db.DbManager()
	connectString := fmt.Sprintf("./storage/biotoken.db")
	fmt.Println("connectString: " + connectString)
	db, err := gorm.Open("sqlite3",connectString)
	if err != nil {
		log.Println("Error Connecting to Database")
	}
	db.Close()

	adminParams := new(Adminlib.AdminLoginParams)

	admin := new(models.Admin)
	exists := db.Where("email = ?", adminParams.Email).First(&admin).RecordNotFound()

	if exists {
		return BadRequestResponse(c,lib.AccountNotExist)
	}

	passwordMatch := lib.CompareHashWithPassword(admin.Password, adminParams.Password)
	
	if passwordMatch != true {
		return BadRequestResponse(c, lib.WrongPassword)
	}
	admin.LastVerified = time.Now()
	token, err := Adminlib.GenerateToken(admin)
	response := Adminlib.AdminLoginResponse{
		FullName:   admin.FullName,
		Email:      admin.Email,
		BioAuth:    true,
		VerifiedAt: time.Time{},
		Active:     true,
		Token: token,
	}
	db.Save(&admin)
	return DataResponse(c, response, http.StatusAccepted)
}
