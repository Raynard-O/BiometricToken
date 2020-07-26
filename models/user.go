package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type User struct {
	gorm.Model
	FullName string `json:"full_name"`
	Email string `json:"email"`
	Password string `json:"password"`
	BioAuth bool	`json:"bio_auth"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Active    bool   `json:"active"`
}
