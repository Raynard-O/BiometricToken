package models

import (
	"github.com/jinzhu/gorm"
	"time"
)


type WhoEnrolled struct {
	AdminFullName string	`json:"admin_full_name"`
	AdminEmail	string	`json:"admin_email"`
	AdminID	uint	`json:"admin_id"`
}

type User struct {
	gorm.Model
	FullName     string    `json:"full_name"`
	Email        string    `json:"email"`
	Password     string    `json:"password"`
	BioAuth      bool      `json:"bio_auth"`
	Active       bool      `json:"active"`
	LastVerified time.Time	`json:"last_verified"`
	AdminFullName string	`json:"admin_full_name"`
	AdminEmail	string	`json:"admin_email"`
	AdminID	uint	`json:"admin_id"`
	AdminEnrolled	WhoEnrolled	`json:"admin_enrolled"`
}

