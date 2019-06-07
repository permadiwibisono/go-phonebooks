package models

type Contact struct {
	PrimaryKey
	UserID       uint `json:"user_id"`
	User         User
	PhoneNumbers []PhoneNumber
	Name         string `json:"name"`
	IsFavorited  bool   `json:"is_favorited"`
	Timestamps
}
