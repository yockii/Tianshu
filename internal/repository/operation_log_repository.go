package repository

import (
	"github.com/yockii/Tianshu/internal/model"
	"github.com/yockii/Tianshu/pkg/db"
)

// OperationLogRepository 操作日志相关数据操作
var OperationLogRepository = &operationLogRepository{}

type operationLogRepository struct{}

func (r *operationLogRepository) Create(log *model.OperationLog) error {
	return db.DB.Create(log).Error
}

func (r *operationLogRepository) List(tenantID uint, offset, limit int) ([]*model.OperationLog, int64, error) {
	var logs []*model.OperationLog
	var total int64
	db.DB.Model(&model.OperationLog{}).Where("tenant_id = ?", tenantID).Count(&total)
	err := db.DB.Where("tenant_id = ?", tenantID).Offset(offset).Limit(limit).Find(&logs).Error
	return logs, total, err
}
