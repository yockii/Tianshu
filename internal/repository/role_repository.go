package repository

import (
	"github.com/yockii/Tianshu/internal/model"
	"github.com/yockii/Tianshu/pkg/db"
)

// RoleRepository 角色相关数据操作
var RoleRepository = &roleRepository{}

type roleRepository struct{}

func (r *roleRepository) GetByID(id uint) (*model.Role, error) {
	var role model.Role
	err := db.DB.First(&role, id).Error
	if err != nil {
		return nil, err
	}
	return &role, nil
}

func (r *roleRepository) GetByName(tenantID uint, name string) (*model.Role, error) {
	var role model.Role
	err := db.DB.Where("tenant_id = ? AND name = ?", tenantID, name).First(&role).Error
	if err != nil {
		return nil, err
	}
	return &role, nil
}

func (r *roleRepository) Create(role *model.Role) error {
	return db.DB.Create(role).Error
}

func (r *roleRepository) Update(role *model.Role) error {
	return db.DB.Save(role).Error
}

func (r *roleRepository) Delete(id uint) error {
	return db.DB.Delete(&model.Role{}, id).Error
}

func (r *roleRepository) List(tenantID uint, offset, limit int) ([]*model.Role, int64, error) {
	var roles []*model.Role
	var total int64
	db.DB.Model(&model.Role{}).Where("tenant_id = ?", tenantID).Count(&total)
	err := db.DB.Where("tenant_id = ?", tenantID).Offset(offset).Limit(limit).Find(&roles).Error
	return roles, total, err
}

// UnsetDefaultRoles 将租户下所有角色的 is_default 设置为 false
func (r *roleRepository) UnsetDefaultRoles(tenantID uint) error {
	return db.DB.Model(&model.Role{}).Where("tenant_id = ?", tenantID).Update("is_default", false).Error
}

// GetDefaultRole 获取租户下被标记为默认的角色
func (r *roleRepository) GetDefaultRole(tenantID uint) (*model.Role, error) {
	var role model.Role
	err := db.DB.Where("tenant_id = ? AND is_default = TRUE", tenantID).First(&role).Error
	if err != nil {
		return nil, err
	}
	return &role, nil
}
