package models

// User of system
type User struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	FirstName string `json:"firstname"`
	Email     string `json:"email"`
}
