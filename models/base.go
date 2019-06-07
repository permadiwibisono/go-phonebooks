package models

import "github.com/jinzhu/gorm"

var db *gorm.DB

func AutoMigrate(myDb *gorm.DB) *gorm.DB {
	db = myDb
	myDb.SingularTable(true)
	myDb.Debug().AutoMigrate(&User{}, &Contact{})
	myDb.Model(&Contact{}).AddForeignKey("user_id", "user(id)", "CASCADE", "CASCADE")
	return myDb
}

func GetDB() *gorm.DB {
	return db
}
