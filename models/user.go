package models

import (
	"os"
	"strings"

	h "go-phonebooks/utils/hash"

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
	Password string `json:"-" gorm:"column:_password"`
	Token    string `json:"_token,omitempty" gorm:"-"`
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
	err := GetDB().Where("email = ?", user.Email).First(&userTemp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, false
	}
	if userTemp.ID != 0 {
		msg := "Email address already in use by another user."
		x, y := errors["email"]
		if !y {
			errors["email"] = []string{msg}
		} else {
			errors["email"] = append(x, msg)
		}
	}
	if len(errors) > 0 {
		return map[string]interface{}{"errors": errors}, false
	}
	return nil, true
}

func (user *User) Save() (uint, bool) {
	isNew := GetDB().NewRecord(user)
	if user.Password != "" {
		hashPassword, _ := h.Encrypt(user.Password)
		user.Password = hashPassword
	}
	if isNew {
		GetDB().Create(&user)
		user.Password = ""
	} else {
		GetDB().Save(&user)
		user.Password = ""
	}
	return user.ID, isNew
}

func (user *User) GenerateToken() {
	tk := &Token{UserID: user.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("jwt_token")))
	user.Token = tokenString
}
