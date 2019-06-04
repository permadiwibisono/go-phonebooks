package models

import (
	"strings"

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

func (user *User) Validate() (map[string]interface{}, bool) {
	errors := make(map[string][]string)
	if !strings.Contains(user.Email, "@") {
		msg := "Email address is required."
		x, y := errors["email"]
		if !y {
			errors["email"] = []string{msg}
		} else {
			errors["email"] = append(x, msg)
		}
	}
	if len(user.Password) < 6 {
		msg := "Password address is required."
		x, y := errors["password"]
		if !y {
			errors["password"] = []string{msg}
		} else {
			errors["password"] = append(x, msg)
		}
	}
	userTemp := &User{}
	err := getDB().Where("email = ?", user.Email).First(&userTemp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, false
	}
	if userTemp.ID != 0 {
		msg := "Email address already in use by another user."
		x, y := errors["email"]
		if !y {
			errors["email"] = []string{}
		} else {
			errors["email"] = append(x, msg)
		}
	}
	if len(errors) > 0 {
		return map[string]interface{}{"errors": errors}, false
	}
	return nil, true
}
