package models

// Room model
type Room struct {
    ID          uint   `gorm:"primary_key"`
    Name        string `gorm:"size:255"`
    Area        float64
    Price       float64
    Status      string `gorm:"size:50"`
    Description string `gorm:"size:255"`
}