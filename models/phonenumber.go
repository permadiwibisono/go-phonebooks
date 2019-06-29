package models

type PhoneNumber struct {
	PrimaryKey
	ContactID uint     `json:"contact_id"`
	Contact   *Contact `json:"contact,omitempty" gorm:"PRELOAD:false"`
	Phone     string   `json:"phone" gorm:"size:80"`
	Timestamps
}
