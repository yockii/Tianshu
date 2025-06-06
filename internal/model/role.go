package model

type Role struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	TenantID    uint   `gorm:"index;not null" json:"tenantId"`
	Name        string `gorm:"size:64;not null" json:"name"`
	Description string `gorm:"size:255" json:"description,omitempty"`
	IsDefault   bool   `gorm:"default:false" json:"isDefault,omitempty"`
	CreatedAt   int64  `gorm:"autoCreateTime:milli" json:"createdAt,omitempty"`
	UpdatedAt   int64  `gorm:"autoUpdateTime:milli" json:"updatedAt,omitempty"`
}

type UserRole struct {
	ID     uint `gorm:"primaryKey" json:"id"`
	UserID uint `gorm:"index;not null" json:"userId"`
	RoleID uint `gorm:"index;not null" json:"roleId"`
}

func init() {
	Models = append(Models, &Role{}, &UserRole{})
}
