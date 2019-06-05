package models

import (
	"fmt"
	"os"

	_ "go-phonebooks/utils/env"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func init() {
	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")
	dbPort := os.Getenv("db_port")
	// dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password) //Build connection string
	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, dbHost, dbPort, dbName)
	fmt.Println(dbURI)

	conn, err := gorm.Open("mysql", dbURI)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Connected with databases...")
	}

	db = conn
	db.SingularTable(true)
	db.Debug().AutoMigrate(&User{}, &Contact{}) //Database migration
}

func GetDB() *gorm.DB {
	return db
}
