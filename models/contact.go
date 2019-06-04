package models

import (
	"github.com/jinzhu/gorm"
)

type Contact struct {
	gorm.Model
	UserID uint `json:"user_id"`
	User   User
	Name   string `json:"name"`
	Phone  string `json:"phone"`
}
