package dao

import (
	"github.com/jinzhu/gorm"
	"room-management/models"
)

type RoleDAO struct {
	DB *gorm.DB
}

// AddRole thêm một Role vào cơ sở dữ liệu
func (dao *RoleDAO) AddRole(role *models.Role) error {
	if err := dao.DB.Create(role).Error; err != nil {
		return err
	}
	return nil
}
