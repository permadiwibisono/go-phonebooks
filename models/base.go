package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type PrimaryKey struct {
	ID uint `json:"id" gorm:"primary_key"`
}

type Pagination struct {
	Data        interface{} `json:"data"`
	From        interface{} `json:"from"`
	To          interface{} `json:"to"`
	CurrentPage int         `json:"current_page"`
	NextPage    interface{} `json:"next_page"`
	PrevPage    interface{} `json:"prev_page"`
	FirstPage   interface{} `json:"first_page"`
	LastPage    interface{} `json:"last_page"`
	TotalPage   int         `json:"total_page"`
	PerPage     int         `json:"per_page"`
	Total       int         `json:"total"`
}

type Timestamps struct {
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"-" sql:"index"`
}

type ScopeFunc func(db *gorm.DB) *gorm.DB

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

func ScopePaginate(page int, perPage int, out interface{}, pagination *Pagination) ScopeFunc {
	return func(db *gorm.DB) *gorm.DB {
		offset := (page - 1) * perPage
		pagination.CurrentPage = page
		pagination.PerPage = perPage
		clone := db.NewScope(out)
		clone.DB().Count(&pagination.Total)
		if pagination.Total > 0 {
			pagination.FirstPage = 1
			if pagination.Total%pagination.PerPage > 0 && pagination.Total/pagination.PerPage > 0 {
				pagination.LastPage = (pagination.Total / pagination.PerPage) + 1
			} else if pagination.Total/pagination.PerPage > 1 {
				pagination.LastPage = pagination.Total / pagination.PerPage
			}
			if pagination.LastPage != nil {
				pagination.TotalPage = pagination.LastPage.(int)
			} else {
				pagination.TotalPage = pagination.FirstPage.(int)
			}
			if pagination.CurrentPage != 1 && pagination.CurrentPage-1 >= 1 && pagination.CurrentPage-1 <= pagination.TotalPage {
				pagination.PrevPage = pagination.CurrentPage - 1
			}
			if pagination.CurrentPage+1 <= pagination.TotalPage {
				pagination.NextPage = pagination.CurrentPage + 1
			}
		}
		var ids []uint
		clone.DB().
			Limit(perPage).
			Offset(offset).
			Pluck(clone.PrimaryField().Name, &ids)
		if len(ids) > 0 {
			pagination.From = offset + 1
			pagination.To = offset + len(ids)
		}
		return db.Limit(perPage).Offset(offset)
	}
}
