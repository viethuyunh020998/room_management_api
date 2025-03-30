package services

import (
	"errors"
	"log"
	"room-management/dao"
	"room-management/models"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	UserDAO *dao.UserDAO
}

// AddUser thêm User mới vào cơ sở dữ liệu
func (service *UserService) AddUser(user *models.User) error {

	user.SetStatus("IN_ACTIVE")
	user.SetUsername(user.Email)

	// Mã hóa mật khẩu trước khi lưu vào cơ sở dữ liệu
	if err := user.HashPassword(); err != nil {
		return err
	}

	return service.UserDAO.AddUser(user)
}

func (service *UserService) UpdateUser(user *models.User) error {
	existingUser, err := service.GetUserByEmail(user.Email)
	if err != nil {
		return err
	}

	// Update fields
	existingUser.FistName = user.FistName
	existingUser.LastName = user.LastName
	existingUser.Age = user.Age
	existingUser.Birthday = user.Birthday
	existingUser.Address = user.Address

	return service.UserDAO.DB.Save(existingUser).Error
}

func (service *UserService) GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	if err := service.UserDAO.DB.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, errors.New("user not found")
	}
	return &user, nil
}

func (service *UserService) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := service.UserDAO.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, errors.New("user not found")
	}
	return &user, nil
}

func (service *UserService) CheckEmailExist(email string) (bool, error) {
	var user models.User
	if err := service.UserDAO.DB.Where("email = ?", email).First(&user).Error; err != nil {
		if err.Error() == "record not found" {
			return false, nil // Email không tồn tại
		}
		return false, err // Lỗi khác
	}
	return true, nil // Email tồn tại
}

func (service *UserService) CheckLogin(email, password string) (*models.User, error) {
	var user models.User
	// Tìm người dùng theo email
	if err := service.UserDAO.DB.Where("email = ?", email).First(&user).Error; err != nil {
		log.Println("User not found:", email) // Log khi không tìm thấy người dùng
		return nil, errors.New("user not found")
	}

	// So sánh mật khẩu đã mã hóa với mật khẩu người dùng nhập vào
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, err // Sai mật khẩu
	}

	// Nếu mật khẩu đúng
	log.Println("Login successful for user:", email)
	return &user, nil // Đăng nhập thành công
}
