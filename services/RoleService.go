package services

import (
	"room-management/dao"
	"room-management/models"
)

type RoleService struct {
	RoleDAO *dao.RoleDAO
}

// AddRole thêm Role mới
func (service *RoleService) AddRole(role *models.Role) error {
	return service.RoleDAO.AddRole(role)
}
