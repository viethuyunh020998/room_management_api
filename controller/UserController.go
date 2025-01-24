package controller

import (
	"net/http"
	"room-management/models"
	"room-management/services"

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
func (controller *UserController) Login(c *gin.Context) {
	var loginData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Kiểm tra đăng nhập
	isValid, err := controller.UserService.CheckLogin(loginData.Username, loginData.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error during login"})
		return
	}

	if isValid {
		c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
	}
}
