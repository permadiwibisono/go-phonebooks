package models

type Contact struct {
	PrimaryKey
	UserID       uint `json:"user_id"`
	User         User
	PhoneNumbers []PhoneNumber `json:"phone_numbers" gorm:"ForeignKey:ContactID"`
	Name         string        `json:"name"`
	IsFavorited  bool          `json:"is_favorited"`
	Timestamps
}
