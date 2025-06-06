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

// CheckUserPermission verifies if a user has a specific permission code via their assigned roles
func (s *relationService) CheckUserPermission(userID uint, code string) (bool, error) {
	// SuperAdmin bypass
	if user, err := UserService.GetByID(userID); err == nil && user != nil && user.IsSuperAdmin {
		return true, nil
	}
	// get user roles
	userRoles, err := s.ListRolesByUser(userID)
	if err != nil {
		return false, err
	}
	// iterate roles
	for _, ur := range userRoles {
		// get permissions for role
		rps, err := s.ListPermissionsByRole(ur.RoleID)
		if err != nil {
			return false, err
		}
		for _, rp := range rps {
			perm, err := PermissionService.GetByID(rp.PermissionID)
			if err != nil {
				continue
			}
			if perm != nil && perm.Code == code {
				return true, nil
			}
		}
	}
	return false, nil
}

// ListUserPermissions returns unique permission codes assigned to a user via roles
func (s *relationService) ListUserPermissions(userID uint) ([]string, error) {
	roles, err := s.ListRolesByUser(userID)
	if err != nil {
		return nil, err
	}
	codeSet := make(map[string]struct{})
	for _, ur := range roles {
		rps, err := s.ListPermissionsByRole(ur.RoleID)
		if err != nil {
			continue
		}
		for _, rp := range rps {
			perm, err := PermissionService.GetByID(rp.PermissionID)
			if err != nil || perm == nil {
				continue
			}
			codeSet[perm.Code] = struct{}{}
		}
	}
	codes := make([]string, 0, len(codeSet))
	for code := range codeSet {
		codes = append(codes, code)
	}
	return codes, nil
}
