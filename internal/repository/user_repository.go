package repository

import (
	"github.com/yockii/Tianshu/internal/model"
	"github.com/yockii/Tianshu/pkg/db"
)

// UserRepository 用户相关数据操作
var UserRepository = &userRepository{}

type userRepository struct{}

func (r *userRepository) GetByID(id uint) (*model.User, error) {
	var user model.User
	err := db.DB.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetByUsername(tenantID uint, username string) (*model.User, error) {
	var user model.User
	err := db.DB.Where("tenant_id = ? AND username = ?", tenantID, username).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) Create(user *model.User) error {
	return db.DB.Create(user).Error
}

func (r *userRepository) Update(user *model.User) error {
	return db.DB.Save(user).Error
}

func (r *userRepository) Delete(id uint) error {
	return db.DB.Delete(&model.User{}, id).Error
}

func (r *userRepository) List(tenantID uint, offset, limit int) ([]*model.User, int64, error) {
	var users []*model.User
	var total int64
	db.DB.Model(&model.User{}).Where("tenant_id = ?", tenantID).Count(&total)
	err := db.DB.Where("tenant_id = ?", tenantID).Offset(offset).Limit(limit).Find(&users).Error
	return users, total, err
}

// 可继续实现RoleRepository、PermissionRepository等...
