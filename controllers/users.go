package controllers

import (
	"BiometricToken/db"
	"BiometricToken/lib"
	Userlib "BiometricToken/lib/user"
	"BiometricToken/models"
	"fmt"
	"github.com/jinzhu/gorm"
	"log"

	//"fmt"
	//"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	//"log"
	"net/http"
	"time"
)

func RegisterUsers(c echo.Context)	error  {
	db := db.DbManager()
	connectString := fmt.Sprintf("./storage/biotoken.db")
	fmt.Println("connectString: " + connectString)
	db, err := gorm.Open("sqlite3",connectString)
	if err != nil {
		log.Println("Error Connecting to Database")
	}
	db.Close()
	params := new(Userlib.RegisterParams)

	if err := c.Bind(params); err !=nil {
		return BadRequestResponse(c,lib.INVALID_BODY)
	}
	log.Println(params)
	user := new(models.User)
	exists := db.Where("email = ?", params.Email).Find(&user).RecordNotFound()
	if exists {
		return BadRequestResponse(c,lib.AccountExists)
	}
	user.Email	=	params.Email
	user.FullName	=	params.FullName
	user.CreatedAt	=	time.Now()
	user.BioAuth	=	params.BioAuth
	user.Active	=	true
	user.Password	=	lib.GenerateHashFromPassword(params.Password)
	db.Create(&user)
	db.Save(&user)
	exists = db.Where("email= ?", user.Email).Find(&user).RecordNotFound()
	if exists == true {
		return BadRequestResponse(c,"Error saving user details")

	}

	return DataResponse(c, user, http.StatusOK)
}

func Verify(c echo.Context)	error {
	//init the db
	db := db.DbManager()
	connectString := fmt.Sprintf("./storage/biotoken.db")
	fmt.Println("connectString: " + connectString)
	db, err := gorm.Open("sqlite3",connectString)
	if err != nil {
		log.Println("Error Connecting to Database")
	}
	db.Close()
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
	user.BioAuth = params.BioAuth	//change this to the input from device
	if user.BioAuth != true	{
		return BadRequestResponse(c,"BioAuth invalid")
	}
	lastVerified := time.Now()
	user.LastVerified = lastVerified
	db.Save(&user)
	//response
	response := Userlib.UserResonse{
		FullName:   user.FullName,
		Email:      user.Email,
		BioAuth:    params.BioAuth,
		VerifiedAt:	lastVerified,
		Active:     true,
	}

	return DataResponse(c,response,http.StatusAccepted)
}



func GetUsers(c echo.Context) error  {
	db := db.DbManager()
	connectString := fmt.Sprintf("./storage/biotoken.db")
	fmt.Println("connectString: " + connectString)
	db, err := gorm.Open("sqlite3",connectString)
	if err != nil {
		log.Println("Error Connecting to Database")
	}
	db.Close()
	var users []models.User

	db.Find(&users)

	return c.JSONPretty(200,users,"")
	//return DataResponse(c, users, http.StatusOK)
}