package model

type Permission struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Code        string `gorm:"size:64;uniqueIndex;not null" json:"code"`
	Description string `gorm:"size:255" json:"description,omitempty"`
	CreatedAt   int64  `gorm:"autoCreateTime:milli" json:"createdAt,omitempty"`
	UpdatedAt   int64  `gorm:"autoUpdateTime:milli" json:"updatedAt,omitempty"`
}

type RolePermission struct {
	ID           uint `gorm:"primaryKey" json:"id"`
	RoleID       uint `gorm:"index;not null" json:"roleId"`
	PermissionID uint `gorm:"index;not null" json:"permissionId"`
}

func init() {
	Models = append(Models, &Permission{}, &RolePermission{})
}
