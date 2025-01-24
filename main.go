package main

import (
	"log"
	"github.com/gin-gonic/gin"
	"room-management/config"  // Import cấu hình DB
	"room-management/dao"
	"room-management/services"
	"room-management/controller"
	"room-management/middlewares"
	"room-management/models"  // Import models để chạy AutoMigrate
)

func main() {
	// Kết nối đến cơ sở dữ liệu
	config.ConnectDB()
	defer config.DB.Close()

	// Tạo tất cả bảng nếu chưa có (sử dụng AutoMigrate cho tất cả các models)
	if err := config.DB.AutoMigrate(&models.Role{}, &models.User{}, &models.Room{}).Error; err != nil {
		log.Fatal("Error creating tables:", err)
	} else {
		log.Println("Tables created successfully!")
	}

	// Khởi tạo các DAO và Service
	roleDAO := &dao.RoleDAO{DB: config.DB}
	userDAO := &dao.UserDAO{DB: config.DB}
	roomDAO := &dao.RoomDAO{DB: config.DB}

	roleService := &services.RoleService{RoleDAO: roleDAO}
	userService := &services.UserService{UserDAO: userDAO}
	roomService := &services.RoomService{RoomDAO: roomDAO}

	// Khởi tạo các controller
	roleController := &controller.RoleController{RoleService: roleService}
	userController := &controller.UserController{UserService: userService}
	roomController := &controller.RoomController{RoomService: roomService}

	// Khởi tạo router
	router := gin.Default()

	// Các route không cần middleware
	router.POST("/add-role", roleController.AddRole)
	router.POST("/add-user", userController.AddUser)

	// Route cần xác thực
	protected := router.Group("/protected")
	protected.Use(middlewares.AuthMiddleware())  // Sử dụng middleware xác thực
	{
		protected.POST("/add-room", roomController.AddRoom)  // API cần xác thực
	}

	// Chạy ứng dụng
	router.Run(":8080")
}
