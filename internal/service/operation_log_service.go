package service

import (
	"github.com/yockii/Tianshu/internal/model"
	"github.com/yockii/Tianshu/internal/repository"
)

// 操作日志相关业务逻辑
var OperationLogService = &operationLogService{}

type operationLogService struct{}

// List 获取操作日志列表
func (s *operationLogService) List(tenantID uint, offset, limit int) ([]*model.OperationLog, int64, error) {
	return repository.OperationLogRepository.List(tenantID, offset, limit)
}

// Create adds a new operation log entry
func (s *operationLogService) Create(log *model.OperationLog) error {
	return repository.OperationLogRepository.Create(log)
}
