package services

import (
	"room-management/dao"
	"room-management/models"
)

type UserService struct {
	UserDAO *dao.UserDAO
}

// AddUser thêm User mới vào cơ sở dữ liệu
func (service *UserService) AddUser(user *models.User) error {
	// Mã hóa mật khẩu trước khi lưu vào cơ sở dữ liệu
	if err := user.HashPassword(); err != nil {
		return err
	}
	return service.UserDAO.AddUser(user)
}

// CheckLogin kiểm tra đăng nhập với mật khẩu đã mã hóa
func (service *UserService) CheckLogin(username, password string) (bool, error) {
	// Lấy người dùng từ cơ sở dữ liệu
	var user models.User
	err := service.UserDAO.GetUserByUsername(&user, username)
	if err != nil {
		return false, err
	}

	// Kiểm tra mật khẩu
	if user.CheckPasswordHash(password) {
		return true, nil
	}
	return false, nil
}
