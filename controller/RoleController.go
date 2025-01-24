package controller

import (
	"github.com/gin-gonic/gin"
	"room-management/services"
	"room-management/models"
	"net/http"
)

type RoleController struct {
	RoleService *services.RoleService
}

// AddRole xử lý yêu cầu thêm Role
func (controller *RoleController) AddRole(c *gin.Context) {
	var role models.Role
	if err := c.ShouldBindJSON(&role); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := controller.RoleService.AddRole(&role); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add role"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Role added successfully"})
}
