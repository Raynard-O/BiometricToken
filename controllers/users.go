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
	db := db.DbManager()

	db, err := gorm.Open("sqlite3", "./storage/database.db")
	if err != nil {
		log.Println("Error Connecting to Database")
	}
	defer db.Close()

	context := c.Get("user").(*jwt.Token)
	claims := context.Claims.(jwt.MapClaims)
	adminEmail := claims["email"]

	adminAuth := new(models.Admin)
	exists := db.Where("email = ?", adminEmail).First(&adminAuth).RecordNotFound()
	if exists {
		return BadRequestResponse(c,"Admin Verification from jwt claims Error")
	}


	params := new(Userlib.RegisterParams)

	if err := c.Bind(params); err !=nil {
		return BadRequestResponse(c,lib.INVALID_BODY)
	}
	log.Println(params)

	user := new(models.User)

	exists = db.Where("email = ?", params.Email).Find(&user).RecordNotFound()
	if exists {
		return BadRequestResponse(c,lib.AccountExists)
	}
	newUser := models.User{

		FullName:      params.FullName,
		Email:         params.Email,
		Password:      lib.GenerateHashFromPassword(params.Password),
		BioAuth:       true,
		CreatedAt:     time.Now(),
		Active:        true,
		AdminEnrolled: models.WhoEnrolled{
			AdminEmail: adminAuth.Email,
			AdminFullName: adminAuth.FullName,
			AdminID: adminAuth.ID,
		},
	}

	db.Save(&newUser)
	exists = db.Where("email= ?", newUser.Email).Find(&user).RecordNotFound()
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

// verify user in db then generate jwt for verify page


//verify user with biodata
func Verify(c echo.Context)	error {
	//init the db
	db := db.DbManager()

	db, err := gorm.Open("sqlite3", "./storage/database.db")
	if err != nil {
		log.Println("Error Connecting to Database")
	}
	defer db.Close()
	//use the email to pull user from db
	params := new(Userlib.VerifyParams)
	if err := c.Bind(params); err != nil {
		panic("error get input data")
	}
	//query db for user
	user := new(models.User)
	exists := db.Where("email = ?", params.Email).Find(&user).RecordNotFound()
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
	db.Save(&user)
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
	db := db.DbManager()

	db, err := gorm.Open("sqlite3", "./storage/database.db")
	if err != nil {
		log.Println("Error Connecting to Database")
	}
	defer db.Close()
	var users []models.User

	db.Find(&users)

	return c.JSONPretty(200,users,"")
	//return DataResponse(c, users, http.StatusOK)
}