package db

import (
	"BiometricToken/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
)

var db *gorm.DB
var err error

func DbInit()  {


	db, err := gorm.Open("sqlite3", "./storage/database.db")
	if err != nil {
		log.Println("Error Connecting to Database")
	}
	defer db.Close()

	db.AutoMigrate(&models.Stats{})
}


func DbManager() *gorm.DB {
	return db
}
