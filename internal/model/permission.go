package model

type Permission struct {
	ID          uint   `gorm:"primaryKey"`
	Code        string `gorm:"size:64;uniqueIndex;not null"`
	Description string `gorm:"size:255"`
	CreatedAt   int64  `gorm:"autoCreateTime:milli"`
	UpdatedAt   int64  `gorm:"autoUpdateTime:milli"`
}

type RolePermission struct {
	ID           uint `gorm:"primaryKey"`
	RoleID       uint `gorm:"index;not null"`
	PermissionID uint `gorm:"index;not null"`
}

func init() {
	Models = append(Models, &Permission{}, &RolePermission{})
}
