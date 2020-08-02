package controllers

import (
	"BiometricToken/db"
	"BiometricToken/lib"
	Userlib "BiometricToken/lib/user"
	"BiometricToken/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
	"time"
)


func RegisterUsers(c echo.Context)	error  {
	Db := db.DbManager()

	Db, err := gorm.Open("postgres", ".user=raynardomongbale password=raynard dbname=biotoken sslmode=disable")
	if err != nil {
		log.Println("Error Connecting to Database")
	}
	defer Db.Close()

	context := c.Get("user").(*jwt.Token)
	claims := context.Claims.(jwt.MapClaims)
	adminEmail := claims["email"]

	adminAuth := new(models.Admin)
	exists := Db.Where("email = ?", adminEmail).First(&adminAuth).RecordNotFound()
	if exists {
		return BadRequestResponse(c,"Admin Verification from jwt claims Error")
	}


	params := new(Userlib.RegisterParams)

	if err := c.Bind(params); err !=nil {
		return BadRequestResponse(c,lib.INVALID_BODY)
	}
	log.Println(params)
	if params.Password !=	params.ConfirmPassword {
		return BadRequestResponse(c, "Input Password Error")
	}

	user := new(models.User)

	exists = Db.Where("email = ?", params.Email).Find(&user).RecordNotFound()
	if !exists {
		return BadRequestResponse(c,lib.AccountExists)
	}
	newUser := models.User{

		FullName:      params.FullName,
		Email:         params.Email,
		Password:      lib.GenerateHashFromPassword(params.Password),
		BioAuth:       true,
		Active:        true,
		AdminEmail: adminAuth.Email,
		AdminFullName: adminAuth.FullName,
		AdminID: adminAuth.ID,
	}

	Db.Save(&newUser)
	exists = Db.Where("email= ?", newUser.Email).Find(&user).RecordNotFound()
	if exists == true {
		return BadRequestResponse(c,"Error saving user details")

	}
	response := Userlib.RegisterResponse{
		User:    Userlib.UserResponse{
			FullName: newUser.FullName,
			Email: newUser.Email,
			BioAuth: newUser.BioAuth,
			Active: newUser.Active,
		},
		Success: true,
		Admin:   Userlib.AdminDetail{
			AdminID: adminAuth.ID,
			AdminName: adminAuth.FullName,
		},
	}
	return DataResponse(c, response, http.StatusOK)
}


//verify user with biodata
func Verify(c echo.Context)	error {
	//init the Db
	Db := db.DbManager()

	Db, err := gorm.Open("postgres", ".user=raynardomongbale password=raynard dbname=biotoken sslmode=disable")
	if err != nil {
		log.Println("Error Connecting to Database")
	}
	defer Db.Close()


	//use the email to pull user from Db
	params := new(Userlib.VerifyParams)
	if err := c.Bind(params); err != nil {
		panic("error get input data")
	}
	//query Db for user
	user := new(models.User)
	exists := Db.Where("email = ?", params.Email).Find(&user).RecordNotFound()
	if exists == true {
		return BadRequestResponse(c,"Error saving user details")
	}
	//confirm user authentication from biodevice
	//change this to the input from device
	if params.BioAuth != true	{
		return BadRequestResponse(c,"BioAuth invalid")
	}
	lastVerified := time.Now()
	user.LastVerified = lastVerified
	Db.Save(&user)
	//response
	response := Userlib.UserVerifyResonse{
		FullName:   user.FullName,
		Email:      user.Email,
		BioAuth:    params.BioAuth,
		VerifiedAt:	lastVerified,
		Active:     true,
	}
	//return c.Redirect(200, "/verify")

	return DataResponse(c,response,http.StatusAccepted)
}



func GetUsers(c echo.Context) error  {
	Db := db.DbManager()

	Db, err := gorm.Open("postgres", ".user=raynardomongbale password=raynard dbname=biotoken sslmode=disable")
	if err != nil {
		log.Println("Error Connecting to Database")
	}
	defer Db.Close()


	var users []models.User

	Db.Find(&users)

	return DataResponse(c, users, http.StatusOK)
}