package models

type Contact struct {
	PrimaryKey
	UserID       uint          `json:"user_id"`
	User         *User         `json:"user,omitempty" gorm:"PRELOAD:false;"`
	PhoneNumbers []PhoneNumber `json:"phone_numbers,omitempty" gorm:"ForeignKey:ContactID;PRELOAD:false"`
	Name         string        `json:"name"`
	IsFavorited  bool          `json:"is_favorited"`
	Timestamps
}
