package models

type PhoneNumber struct {
	PrimaryKey
	ContactID uint `json:"contact_id"`
	Contact   Contact
	Phone     string `json:"phone" gorm:"size:80"`
	Timestamps
}
