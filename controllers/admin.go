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
	db := db.DbManager()

	db, err := gorm.Open("postgres", ".user=raynardomongbale password=raynard dbname=biometrictoken sslmode=disable")
	if err != nil {
		log.Println("Error Connecting to Database")
	}
	defer db.Close()

	adminParams := new(Adminlib.AdminLoginParams)

	admin := new(models.Admin)
	exists := db.Where("email = ?", adminParams.Email).First(&admin).RecordNotFound()

	if exists {
		return BadRequestResponse(c,lib.AccountNotExist)
	}
	if admin.ID == 0 {
		return BadRequestResponse(c, lib.AccountNotExist)
	}


	passwordMatch := lib.CompareHashWithPassword(admin.Password, adminParams.Password)
	
	if !passwordMatch  || !adminParams.Auth {
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

func CreateAdmin(c echo.Context)	error	{
	db := db.DbManager()

	db, err := gorm.Open("postgres", ".user=raynardomongbale password=raynard dbname=biometrictoken sslmode=disable")
	if err != nil {
		log.Println("Error Connecting to Database")
	}
	defer db.Close()


	params := new(Adminlib.CreateAdminParams)
	if err := c.Bind(params); err != nil {
		panic("error getting params")
	}
	admin := new(models.Admin)
	exists := db.Where("email = ?", params.Email).First(&admin).RecordNotFound()

	if !exists {
		return BadRequestResponse(c,lib.AccountExists)
	}
	admin.FullName = params.FullName
	admin.Email = params.Email
	admin.Password = lib.GenerateHashFromPassword(params.Password)
	admin.CreatedAt = time.Now()
	admin.BioAuth = true
	admin.Active = true
	db.Create(&admin)
	db.Save(&admin)

	adminResponse := Adminlib.AdminCreateResponse{
		FullName: admin.FullName,
		Email:    admin.Email,
		BioAuth:  admin.BioAuth,
		Active:   admin.Active,
	}
	exists = db.Where("email= ?", admin.Email).Find(&admin).RecordNotFound()
	if exists == true {
		return c.JSON(http.StatusNotModified, lib.AccountNotExist)
	}


	return DataResponse(c, adminResponse, http.StatusAccepted)
}


func GetAdmin(c echo.Context) error  {
	db := db.DbManager()

	db, err := gorm.Open("postgres", ".user=raynardomongbale password=raynard dbname=biometrictoken sslmode=disable")
	if err != nil {
		log.Println("Error Connecting to Database")
	}
	defer db.Close()

	var admins []models.Admin

	db.Find(&admins)


	return DataResponse(c, admins, http.StatusOK)
}