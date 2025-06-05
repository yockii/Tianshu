package service

import (
	"github.com/yockii/Tianshu/internal/model"
	"github.com/yockii/Tianshu/internal/repository"
)

// 权限相关业务逻辑
var PermissionService = &permissionService{}

type permissionService struct{}

func (s *permissionService) GetByID(id uint) (*model.Permission, error) {
	return repository.PermissionRepository.GetByID(id)
}

func (s *permissionService) GetByCode(code string) (*model.Permission, error) {
	return repository.PermissionRepository.GetByCode(code)
}

func (s *permissionService) Create(perm *model.Permission) error {
	return repository.PermissionRepository.Create(perm)
}

func (s *permissionService) Update(perm *model.Permission) error {
	return repository.PermissionRepository.Update(perm)
}

func (s *permissionService) List(offset, limit int) ([]*model.Permission, int64, error) {
	return repository.PermissionRepository.List(offset, limit)
}
