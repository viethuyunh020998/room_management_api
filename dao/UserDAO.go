package dao

import (
	"github.com/jinzhu/gorm"
	"room-management/models"
)

type UserDAO struct {
	DB *gorm.DB
}

// AddUser thêm người dùng vào cơ sở dữ liệu
func (dao *UserDAO) AddUser(user *models.User) error {
	if err := dao.DB.Create(user).Error; err != nil {
		return err
	}
	return nil
}


// GetUserByUsername lấy người dùng từ tên đăng nhập
func (dao *UserDAO) GetUserByUsername(user *models.User, username string) error {
	if err := dao.DB.Where("username = ?", username).First(user).Error; err != nil {
		return err
	}
	return nil
}