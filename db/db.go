package db

import (
	"BiometricToken/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

var db *gorm.DB
var err error

func DbInit()  {


	db, err := gorm.Open("postgres", ".user=raynardomongbale password=raynard dbname=biotoken sslmode=disable")
	if err != nil {
		log.Println("Error Connecting to Database")
	}
	defer db.Close()

	db.AutoMigrate(&models.Admin{},&models.User{})
}


func DbManager() *gorm.DB {
	return db
}
