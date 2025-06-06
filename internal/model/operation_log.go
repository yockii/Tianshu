package model

type OperationLog struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	TenantID  uint   `gorm:"index;not null" json:"tenantId"`
	UserID    uint   `gorm:"index;not null" json:"userId"`
	Action    string `gorm:"size:128;not null" json:"action"`
	Detail    string `gorm:"type:text" json:"detail,omitempty"`
	CreatedAt int64  `gorm:"autoCreateTime:milli" json:"createdAt"`
}

func init() {
	Models = append(Models, &OperationLog{})
}
