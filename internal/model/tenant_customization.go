package model

type TenantCustomization struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	TenantID    uint   `gorm:"index;not null" json:"tenantId"`
	Logo        string `gorm:"size:255" json:"logo,omitempty"`
	SiteName    string `gorm:"size:64" json:"siteName,omitempty"`
	ThemeColor  string `gorm:"size:32" json:"themeColor,omitempty"`
	Favicon     string `gorm:"size:255" json:"favicon,omitempty"`
	ExtraConfig string `gorm:"type:jsonb" json:"extraConfig,omitempty"`
	CreatedAt   int64  `gorm:"autoCreateTime:milli" json:"createdAt"`
	UpdatedAt   int64  `gorm:"autoUpdateTime:milli" json:"updatedAt,omitempty"`
}

func init() {
	Models = append(Models, &TenantCustomization{})
}
