package service

import (
	"github.com/yockii/Tianshu/internal/model"
	"github.com/yockii/Tianshu/internal/repository"
)

// 角色-权限、用户-角色分配相关业务逻辑
var RelationService = &relationService{}

type relationService struct{}

// 角色分配权限
func (s *relationService) AssignPermissionToRole(roleID, permID uint) error {
	return repository.RolePermissionRepository.Assign(roleID, permID)
}

func (s *relationService) RemovePermissionFromRole(roleID, permID uint) error {
	return repository.RolePermissionRepository.Remove(roleID, permID)
}

func (s *relationService) ListPermissionsByRole(roleID uint) ([]*model.RolePermission, error) {
	return repository.RolePermissionRepository.ListByRole(roleID)
}

// 用户分配角色
func (s *relationService) AssignRoleToUser(userID, roleID uint) error {
	return repository.UserRoleRepository.Assign(userID, roleID)
}

func (s *relationService) RemoveRoleFromUser(userID, roleID uint) error {
	return repository.UserRoleRepository.Remove(userID, roleID)
}

func (s *relationService) ListRolesByUser(userID uint) ([]*model.UserRole, error) {
	return repository.UserRoleRepository.ListByUser(userID)
}
