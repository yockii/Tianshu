package repository

import (
	"github.com/yockii/Tianshu/internal/model"
	"github.com/yockii/Tianshu/pkg/db"
)

// PermissionRepository 权限相关数据操作
var PermissionRepository = &permissionRepository{}

type permissionRepository struct{}

func (r *permissionRepository) GetByID(id uint) (*model.Permission, error) {
	var perm model.Permission
	err := db.DB.First(&perm, id).Error
	if err != nil {
		return nil, err
	}
	return &perm, nil
}

func (r *permissionRepository) GetByCode(code string) (*model.Permission, error) {
	var perm model.Permission
	err := db.DB.Where("code = ?", code).First(&perm).Error
	if err != nil {
		return nil, err
	}
	return &perm, nil
}

func (r *permissionRepository) Create(perm *model.Permission) error {
	return db.DB.Create(perm).Error
}

func (r *permissionRepository) Update(perm *model.Permission) error {
	return db.DB.Save(perm).Error
}

func (r *permissionRepository) Delete(id uint) error {
	return db.DB.Delete(&model.Permission{}, id).Error
}

func (r *permissionRepository) List(offset, limit int) ([]*model.Permission, int64, error) {
	var perms []*model.Permission
	var total int64
	db.DB.Model(&model.Permission{}).Count(&total)
	err := db.DB.Offset(offset).Limit(limit).Find(&perms).Error
	return perms, total, err
}
