package models

import (
	"time"

	"crypto/rand"
	"encoding/hex"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// User model
type User struct {
	ID           uint   `gorm:"primary_key;autoIncrement"`
	Username     string `gorm:"size:255;unique"`
	Password     string `gorm:"size:255"`
	Email        string `gorm:"size:255;unique"`
	FistName     string `gorm:"size:20"`
	LastName     string `gorm:"size:20"`
	Age          int    `gorm:"index"`
	IsVerified   bool
	Address      string `gorm:"size:255"`
	Birthday     time.Time
	Status       string    `gorm:"size:50"`
	CreatedDate  time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	ModifiedDate time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	RoleID       int       `gorm:"index"`
}

// Method to set status of user
func (u *User) SetStatus(status string) {
	u.Status = status
}

func (u *User) SetUsername(username string) {
	u.Username = username
}

// HashPassword mã hóa mật khẩu
func (user *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return nil
}

// CheckPasswordHash kiểm tra mật khẩu khi đăng nhập
func (user *User) CheckPasswordHash(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}

func GenerateToke() (string, error) {
	token := make([]byte, 16) // Tạo token 16 byte
	_, err := rand.Read(token)
	if err != nil {
		return "", fmt.Errorf("error generating token: %v", err)
	}
	return hex.EncodeToString(token), nil

}
