package config

import (
	"log"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB
var err error

func ConnectDB() {
	// B1: Kết nối MySQL không chọn database trước
	tempDB, err := gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal("Initial connection failed: ", err)
	}
	defer tempDB.Close()

	// B2: Tạo database nếu chưa có
	tempDB.Exec("CREATE DATABASE IF NOT EXISTS room_management")

	// B3: Kết nối lại với DB chính
	DB, err = gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/room_management?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal("Connection to room_management failed: ", err)
	}

	log.Println("✅ Connected to the database")
}
