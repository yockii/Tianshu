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

// 删除角色
func (s *roleService) Delete(id uint) error {
	return repository.RoleRepository.Delete(id)
}

// UnsetDefaultRoles 清除租户下所有角色的默认标记
func (s *roleService) UnsetDefaultRoles(tenantID uint) error {
	return repository.RoleRepository.UnsetDefaultRoles(tenantID)
}

// GetDefaultRole 获取租户下的默认角色
func (s *roleService) GetDefaultRole(tenantID uint) (*model.Role, error) {
	return repository.RoleRepository.GetDefaultRole(tenantID)
}
