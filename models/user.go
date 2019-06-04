package models

import (
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
)

type Token struct {
	UserID uint
	jwt.StandardClaims
}

type User struct {
	gorm.Model
	Email    string `json:"email"`
	Password string `json:"password" gorm:"column:_password"`
	Token    string `json:"_token" sql:"-"`
}
