package dao

import (
	"github.com/jinzhu/gorm"
	"room-management/models"
)

type RoomDAO struct {
	DB *gorm.DB
}

func (dao *RoomDAO) AddRoom(room *models.Room) error {
	if err := dao.DB.Create(room).Error; err != nil {
		return err
	}
	return nil
}

func (dao *RoomDAO) GetRoomByID(id uint) (*models.Room, error) {
	var room models.Room
	if err := dao.DB.First(&room, id).Error; err != nil {
		return nil, err
	}
	return &room, nil
}
