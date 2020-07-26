package db

import (
	model "BiometricToken/models"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func DbManager() *gorm.DB {
	return db
}

func DbInit()  {
	db, err := gorm.Open("sqlite3","./storage/db.db")
	if err != nil {
		panic("Error Connecting to Database")
	}
	defer db.Close()

	db.AutoMigrate(&model.Stats{})
}