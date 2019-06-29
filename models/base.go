package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type PrimaryKey struct {
	ID uint `json:"id" gorm:"primary_key"`
}

type Timestamps struct {
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"-" sql:"index"`
}

func AutoMigrate(myDb *gorm.DB) *gorm.DB {
	db = myDb
	myDb.LogMode(true)
	myDb.SingularTable(true)
	myDb.Debug().AutoMigrate(&User{}, &Contact{}, &PhoneNumber{})
	myDb.Debug().Model(&Contact{}).AddForeignKey("user_id", "user(id)", "CASCADE", "CASCADE")
	myDb.Debug().Model(&PhoneNumber{}).AddForeignKey("contact_id", "contact(id)", "CASCADE", "CASCADE")
	return myDb
}

func GetDB() *gorm.DB {
	return db
}
