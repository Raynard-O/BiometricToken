package models

import (
	"github.com/jinzhu/gorm"
	"time"
)


type Admin struct {
	gorm.Model
	FullName     string    `json:"full_name"`
	Email        string    `json:"email"`
	Password     string    `json:"password"`
	BioAuth      bool      `json:"bio_auth"`
	Active       bool      `json:"active"`
	LastVerified time.Time	`json:"last_verified"`
}



