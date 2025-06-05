package service

import (
	"github.com/yockii/Tianshu/internal/model"
	"github.com/yockii/Tianshu/internal/repository"
)

// 用户相关业务逻辑
var UserService = &userService{}

type userService struct{}

func (s *userService) GetByID(id uint) (*model.User, error) {
	return repository.UserRepository.GetByID(id)
}

func (s *userService) GetByUsername(tenantID uint, username string) (*model.User, error) {
	return repository.UserRepository.GetByUsername(tenantID, username)
}

func (s *userService) Create(user *model.User) error {
	return repository.UserRepository.Create(user)
}

func (s *userService) Update(user *model.User) error {
	return repository.UserRepository.Update(user)
}

func (s *userService) List(tenantID uint, offset, limit int) ([]*model.User, int64, error) {
	return repository.UserRepository.List(tenantID, offset, limit)
}
