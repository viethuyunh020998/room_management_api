package models

import (
	"golang.org/x/crypto/bcrypt"
	"time"
)

// User model
type User struct {
	ID          uint      `gorm:"primary_key"`
	Username    string    `gorm:"size:255;unique"`
	Password    string    `gorm:"size:255"` // Mã hóa mật khẩu trước khi lưu
	Email       string    `gorm:"size:255;unique"`
	FullName    string    `gorm:"size:255"`
	CreatedDate time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	ModifiedDate time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	RoleID      int       `gorm:"index"`
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
