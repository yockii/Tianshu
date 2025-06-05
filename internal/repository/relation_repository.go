package repository

import (
	"github.com/yockii/Tianshu/internal/model"
	"github.com/yockii/Tianshu/pkg/db"
)

// RolePermissionRepository 角色权限分配相关数据操作
var RolePermissionRepository = &rolePermissionRepository{}

type rolePermissionRepository struct{}

func (r *rolePermissionRepository) Assign(roleID, permID uint) error {
	return db.DB.Create(&model.RolePermission{RoleID: roleID, PermissionID: permID}).Error
}

func (r *rolePermissionRepository) Remove(roleID, permID uint) error {
	return db.DB.Where("role_id = ? AND permission_id = ?", roleID, permID).Delete(&model.RolePermission{}).Error
}

func (r *rolePermissionRepository) ListByRole(roleID uint) ([]*model.RolePermission, error) {
	var list []*model.RolePermission
	err := db.DB.Where("role_id = ?", roleID).Find(&list).Error
	return list, err
}

// UserRoleRepository 用户角色分配相关数据操作
var UserRoleRepository = &userRoleRepository{}

type userRoleRepository struct{}

func (r *userRoleRepository) Assign(userID, roleID uint) error {
	return db.DB.Create(&model.UserRole{UserID: userID, RoleID: roleID}).Error
}

func (r *userRoleRepository) Remove(userID, roleID uint) error {
	return db.DB.Where("user_id = ? AND role_id = ?", userID, roleID).Delete(&model.UserRole{}).Error
}

func (r *userRoleRepository) ListByUser(userID uint) ([]*model.UserRole, error) {
	var list []*model.UserRole
	err := db.DB.Where("user_id = ?", userID).Find(&list).Error
	return list, err
}
