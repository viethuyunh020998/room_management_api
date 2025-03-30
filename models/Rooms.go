package models

import (
	"time"
)

// Room model
type Room struct {
    ID          uint   `gorm:"primary_key;autoIncrement"`
    Name        string `gorm:"size:255"`
    Area        float64
    Price       float64
    Status      string `gorm:"size:50"`
    Description string `gorm:"size:255"`
    CreatedDate      time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	CreatedByUserID  int
	ModifiedDate     time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	ModifiedByUserID int `gorm:"index"`
}