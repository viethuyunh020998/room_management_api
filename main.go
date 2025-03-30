package main

import (
	"log"
	"room-management/config" // Import cấu hình DB
	"room-management/controller"
	"room-management/dao"
	"room-management/middlewares"
	"room-management/models" // Import models để chạy AutoMigrate
	"room-management/services"

	"github.com/gin-contrib/cors" // Import thư viện CORS
	"github.com/gin-gonic/gin"
	// "github.com/joho/godotenv"
	// "golang.org/x/crypto/bcrypt"
)

func main() {
	// if err := godotenv.Load(); err != nil {
	// 	log.Println("⚠️ Không tìm thấy file .env, dùng giá trị mặc định nếu có.")
	// }

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

	// Cấu hình CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},                   // Cho phép frontend từ localhost:3000
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},            // Các phương thức HTTP được phép
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"}, // Các headers được phép
		AllowCredentials: true,                                                // Cho phép gửi cookie
	}))

	// Các route không cần middleware
	router.POST("/add-role", roleController.AddRole)
	router.POST("/add-user", userController.AddUser)
	router.POST("/login", userController.Login)
	router.GET("/check-email", userController.CheckEmail)
	// Route cần xác thực
	protected := router.Group("/protected")
	protected.Use(middlewares.AuthMiddleware()) // Sử dụng middleware xác thực
	{
		// API need Authenticate
		protected.POST("/add-room", roomController.AddRoom)
		protected.POST("/update-user", userController.EditUser)
	}
	// Load variables from .env file
	// if err := godotenv.Load(); err != nil {
	// 	log.Fatal("Error loading .env file")
	// }

	// Chạy ứng dụng
	router.Run(":8080")

}
