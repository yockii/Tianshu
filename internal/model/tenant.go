package model

type Tenant struct {
	ID            uint                `gorm:"primaryKey" json:"id"`
	Name          string              `gorm:"size:64;not null;uniqueIndex" json:"name"`
	Logo          string              `gorm:"size:255" json:"logo,omitempty"`
	Theme         string              `gorm:"size:32" json:"theme,omitempty"`
	Domain        string              `gorm:"size:128;not null;uniqueIndex" json:"domain"`
	WelcomeText   string              `gorm:"size:255" json:"welcomeText,omitempty"`
	Customization TenantCustomization `gorm:"foreignKey:TenantID" json:"customization,omitempty"`
	CreatedAt     int64               `gorm:"autoCreateTime:milli" json:"createdAt"`
	UpdatedAt     int64               `gorm:"autoUpdateTime:milli" json:"updatedAt,omitempty"`
	DeletedAt     int64               `gorm:"index" json:"deletedAt,omitempty"`
}

func init() {
	Models = append(Models, &Tenant{})
}
