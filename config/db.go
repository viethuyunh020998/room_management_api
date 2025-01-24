package config

import (
	"log"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// DB lưu trữ kết nối cơ sở dữ liệu
var DB *gorm.DB
var err error

// ConnectDB thiết lập kết nối đến cơ sở dữ liệu
func ConnectDB() {
	DB, err = gorm.Open("mysql", "root:viet123@/room_management?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal("Connection failed: ", err)
	}
	log.Println("Connected to the database")
}
