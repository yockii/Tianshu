package repository

import (
	"github.com/yockii/Tianshu/internal/model"
	"github.com/yockii/Tianshu/pkg/db"
)

// TenantRepository 租户相关数据操作
var TenantRepository = &tenantRepository{}

type tenantRepository struct{}

func (r *tenantRepository) GetByID(id uint) (*model.Tenant, error) {
	var tenant model.Tenant
	err := db.DB.Preload("Customization").First(&tenant, id).Error
	if err != nil {
		return nil, err
	}
	return &tenant, nil
}

func (r *tenantRepository) GetByDomain(domain string) (*model.Tenant, error) {
	var tenant model.Tenant
	err := db.DB.Preload("Customization").Where("domain = ?", domain).First(&tenant).Error
	if err != nil {
		return nil, err
	}
	return &tenant, nil
}

func (r *tenantRepository) Create(tenant *model.Tenant) error {
	return db.DB.Create(tenant).Error
}

func (r *tenantRepository) Update(tenant *model.Tenant) error {
	return db.DB.Save(tenant).Error
}

func (r *tenantRepository) Delete(id uint) error {
	return db.DB.Delete(&model.Tenant{}, id).Error
}

func (r *tenantRepository) List(offset, limit int) ([]*model.Tenant, int64, error) {
	var tenants []*model.Tenant
	var total int64
	db.DB.Model(&model.Tenant{}).Count(&total)
	err := db.DB.Preload("Customization").Offset(offset).Limit(limit).Find(&tenants).Error
	return tenants, total, err
}
