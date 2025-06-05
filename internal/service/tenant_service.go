package service

import (
	"github.com/yockii/Tianshu/internal/model"
	"github.com/yockii/Tianshu/internal/repository"
)

// 租户相关业务逻辑
var TenantService = &tenantService{}

type tenantService struct{}

func (s *tenantService) GetByID(id uint) (*model.Tenant, error) {
	return repository.TenantRepository.GetByID(id)
}

func (s *tenantService) GetByDomain(domain string) (*model.Tenant, error) {
	return repository.TenantRepository.GetByDomain(domain)
}

func (s *tenantService) Create(tenant *model.Tenant) error {
	return repository.TenantRepository.Create(tenant)
}

func (s *tenantService) Update(tenant *model.Tenant) error {
	return repository.TenantRepository.Update(tenant)
}

func (s *tenantService) List(offset, limit int) ([]*model.Tenant, int64, error) {
	return repository.TenantRepository.List(offset, limit)
}

// 租户定制化内容相关业务逻辑
var TenantCustomizationService = &tenantCustomizationService{}

type tenantCustomizationService struct{}

func (s *tenantCustomizationService) GetByTenantID(tenantID uint) (*model.TenantCustomization, error) {
	return repository.TenantCustomizationRepository.GetByTenantID(tenantID)
}

func (s *tenantCustomizationService) Create(tc *model.TenantCustomization) error {
	return repository.TenantCustomizationRepository.Create(tc)
}

func (s *tenantCustomizationService) Update(tc *model.TenantCustomization) error {
	return repository.TenantCustomizationRepository.Update(tc)
}
