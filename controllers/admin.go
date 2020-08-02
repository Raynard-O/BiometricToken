package controllers

import (
	"BiometricToken/db"
	"BiometricToken/lib"
	Adminlib "BiometricToken/lib/admin"
	"BiometricToken/models"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"time"
)

func LoginAdmin(c echo.Context) error {
	Db := db.DbManager()

	Db, err := gorm.Open("postgres", ".user=raynardomongbale password=raynard dbname=biotoken sslmode=disable")
	if err != nil {
		log.Println("Error Connecting to Database")
	}
	defer Db.Close()

	adminParams := new(Adminlib.AdminLoginParams)
	if err := c.Bind(adminParams); err != nil {
		panic("error binding params")
	}
	admin := new(models.Admin)
	exists := Db.Where("email = ?", adminParams.Email).Find(&admin).RecordNotFound()

	if exists {
		return BadRequestResponse(c,lib.AccountNotExist)
	}
	if admin.ID == 0 {
		return BadRequestResponse(c, lib.AccountNotExist)
	}


	passwordMatch := lib.CompareHashWithPassword(admin.Password, adminParams.Password)
	
	if !passwordMatch  || !adminParams.BioAuth {
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
	Db.Save(&admin)
	return DataResponse(c, response, http.StatusAccepted)
}

func CreateAdmin(c echo.Context)	error	{
	Db := db.DbManager()

	Db, err := gorm.Open("postgres", ".user=raynardomongbale password=raynard dbname=biotoken sslmode=disable")
	if err != nil {
		log.Println("Error Connecting to Database")
	}
	defer Db.Close()


	params := new(Adminlib.CreateAdminParams)
	if err := c.Bind(params); err != nil {
		panic("error getting params")
	}
	if params.Password !=	params.ConfirmPassword {
		return BadRequestResponse(c, "Input Password Error")
	}

	admin := new(models.Admin)
	exists := Db.Where("email = ?", params.Email).First(&admin).RecordNotFound()

	if !exists {
		return BadRequestResponse(c,lib.AccountExists)
	}
	admin.FullName = params.FullName
	admin.Email = params.Email
	admin.Password = lib.GenerateHashFromPassword(params.Password)
	admin.CreatedAt = time.Now()
	admin.BioAuth = params.BioAuth
	admin.Active = true
	Db.Create(&admin)
	Db.Save(&admin)

	adminResponse := Adminlib.AdminCreateResponse{
		FullName: admin.FullName,
		Email:    admin.Email,
		BioAuth:  admin.BioAuth,
		Active:   admin.Active,
	}
	exists = Db.Where("email= ?", admin.Email).Find(&admin).RecordNotFound()
	if exists == true {
		return c.JSON(http.StatusNotModified, lib.AccountNotExist)
	}


	return DataResponse(c, adminResponse, http.StatusAccepted)
}

func GetAdmin(c echo.Context) error  {
	Db := db.DbManager()

	Db, err := gorm.Open("postgres", ".user=raynardomongbale password=raynard dbname=biotoken sslmode=disable")
	if err != nil {
		log.Println("Error Connecting to Database")
	}
	defer Db.Close()

	var admins []models.Admin

	Db.Find(&admins)


	return DataResponse(c, admins, http.StatusOK)
}