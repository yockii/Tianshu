package service

import (
	"github.com/yockii/Tianshu/internal/model"
	"github.com/yockii/Tianshu/internal/repository"
)

// 角色相关业务逻辑
var RoleService = &roleService{}

type roleService struct{}

func (s *roleService) GetByID(id uint) (*model.Role, error) {
	return repository.RoleRepository.GetByID(id)
}

func (s *roleService) GetByName(tenantID uint, name string) (*model.Role, error) {
	return repository.RoleRepository.GetByName(tenantID, name)
}

func (s *roleService) Create(role *model.Role) error {
	return repository.RoleRepository.Create(role)
}

func (s *roleService) Update(role *model.Role) error {
	return repository.RoleRepository.Update(role)
}

func (s *roleService) List(tenantID uint, offset, limit int) ([]*model.Role, int64, error) {
	return repository.RoleRepository.List(tenantID, offset, limit)
}
