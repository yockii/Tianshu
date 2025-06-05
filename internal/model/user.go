package model

type User struct {
	ID           uint   `gorm:"primaryKey" json:"id"`
	TenantID     uint   `gorm:"index;not null" json:"tenantId"`
	Username     string `gorm:"size:64;not null" json:"username"`
	Email        string `gorm:"size:128" json:"email,omitempty"`
	Phone        string `gorm:"size:32" json:"phone,omitempty"`
	PasswordHash string `gorm:"size:255;not null" json:"-"`
	Status       int    `gorm:"default:1" json:"status"`
	IsSuperAdmin bool   `gorm:"default:false" json:"isSuperAdmin"`
	CreatedAt    int64  `gorm:"autoCreateTime:milli" json:"createdAt"`
	UpdatedAt    int64  `gorm:"autoUpdateTime:milli" json:"updatedAt"`
	DeletedAt    int64  `gorm:"index" json:"deletedAt,omitempty"`
}

func init() {
	Models = append(Models, &User{})
}
