package services

import (
	"room-management/dao"
	"room-management/models"
)

type RoomService struct {
	RoomDAO *dao.RoomDAO
}

func (service *RoomService) AddRoom(room *models.Room) error {
	return service.RoomDAO.AddRoom(room)
}
