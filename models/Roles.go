package models

// User model
type Role struct {
	ID       uint   `gorm:"primary_key;autoIncrement"`
	Rolename string `gorm:"size:255;unique"`
}
