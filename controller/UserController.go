package controller

import (
	"log"
	"net/http"
	"os"
	"room-management/models"
	"room-management/services"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// UserController chứa các API liên quan đến User
type UserController struct {
	UserService *services.UserService
}

// AddUser xử lý yêu cầu thêm User
func (controller *UserController) AddUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := controller.UserService.AddUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add user"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User added successfully"})
}

// Login xử lý yêu cầu đăng nhập
// Login xử lý yêu cầu đăng nhập
func (controller *UserController) Login(c *gin.Context) {
	var loginData struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// Bind dữ liệu JSON từ request vào loginData
	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Log thông tin email và mật khẩu sau khi bind dữ liệu
	log.Println("Login attempt: ", loginData.Email)
	log.Println("Password entered: ", loginData.Password)

	// Kiểm tra thông tin đăng nhập
	user, err := controller.UserService.CheckLogin(loginData.Email, loginData.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error during login"})
		return
	}

	if user != nil {
		// Nếu đăng nhập thành công, tạoi JWT token
		token, err := generateJWT(user.ID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generating JWT"})
			return
		}

		// Trả về thông tin người dùng và JWT token
		c.JSON(http.StatusOK, gin.H{
			"message": "Login successful",
			"token":   token,
			"user":    user,
		})
	} else {
		// Nếu thông tin đăng nhập không hợp lệ
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
	}
}

// generateJWT tạo một JWT token cho người dùng
func generateJWT(userID uint) (string, error) {
	// Tạo một JWT token với payload là userID
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 72).Unix(), // Token hết hạn sau 72 giờ
	}

	// Tạo JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Ký token với secret key (bảo mật)
	secretKey := os.Getenv("JWT_SECRET_KEY") // Đảm bảo rằng bạn đã cấu hình secret key trong môi trường
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// CheckEmail xử lý yêu cầu kiểm tra email
func (controller *UserController) CheckEmail(c *gin.Context) {
	email := c.DefaultQuery("email", "")

	if email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email is required"})
		return
	}

	exists, err := controller.UserService.CheckEmailExist(email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error checking email"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"exists": exists})
}

func (controller *UserController) EditUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := controller.UserService.UpdateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add room"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User Update successfully"})
}
