package controller

import (
	"github.com/gin-gonic/gin"
	"room-management/services"
	"room-management/models"
	"net/http"
)

// RoomController chứa các API liên quan đến Room
type RoomController struct {
	RoomService *services.RoomService
}

// AddRoom xử lý yêu cầu thêm Room
func (controller *RoomController) AddRoom(c *gin.Context) {
	var room models.Room
	if err := c.ShouldBindJSON(&room); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := controller.RoomService.AddRoom(&room); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add room"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Room added successfully"})
}
