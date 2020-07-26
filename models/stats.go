package models

import "github.com/jinzhu/gorm"

type Stats struct {
	gorm.Model
	Type string `json:"opt"`
}
