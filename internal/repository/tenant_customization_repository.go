package repository

import (
	"github.com/yockii/Tianshu/internal/model"
	"github.com/yockii/Tianshu/pkg/db"
)

// TenantCustomizationRepository 租户定制化内容相关数据操作
var TenantCustomizationRepository = &tenantCustomizationRepository{}

type tenantCustomizationRepository struct{}

func (r *tenantCustomizationRepository) GetByTenantID(tenantID uint) (*model.TenantCustomization, error) {
	var tc model.TenantCustomization
	err := db.DB.Where("tenant_id = ?", tenantID).First(&tc).Error
	if err != nil {
		return nil, err
	}
	return &tc, nil
}

func (r *tenantCustomizationRepository) Create(tc *model.TenantCustomization) error {
	return db.DB.Create(tc).Error
}

func (r *tenantCustomizationRepository) Update(tc *model.TenantCustomization) error {
	return db.DB.Save(tc).Error
}

func (r *tenantCustomizationRepository) Delete(id uint) error {
	return db.DB.Delete(&model.TenantCustomization{}, id).Error
}
