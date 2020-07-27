package db

import (
	"BiometricToken/models"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
)

var db *gorm.DB
var err error

func DbInit()  {


	connectString := fmt.Sprintf("./Storage/biotoken.db")
	fmt.Println("connectString: " + connectString)
	db, err := gorm.Open("sqlite3",connectString)
	if err != nil {
		log.Println("Error Connecting to Database")
	}
	defer db.Close()

	db.AutoMigrate(&models.Stats{})
}


func DbManager() *gorm.DB {
	return db
}
