package model

type Role struct {
	ID          uint   `gorm:"primaryKey"`
	TenantID    uint   `gorm:"index;not null"`
	Name        string `gorm:"size:64;not null"`
	Description string `gorm:"size:255"`
	IsDefault   bool   `gorm:"default:false"`
	CreatedAt   int64  `gorm:"autoCreateTime:milli"`
	UpdatedAt   int64  `gorm:"autoUpdateTime:milli"`
}

type UserRole struct {
	ID     uint `gorm:"primaryKey"`
	UserID uint `gorm:"index;not null"`
	RoleID uint `gorm:"index;not null"`
}

func init() {
	Models = append(Models, &Role{}, &UserRole{})
}
