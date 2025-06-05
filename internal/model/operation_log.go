package model

type OperationLog struct {
	ID        uint   `gorm:"primaryKey"`
	TenantID  uint   `gorm:"index;not null"`
	UserID    uint   `gorm:"index;not null"`
	Action    string `gorm:"size:128;not null"`
	Detail    string `gorm:"type:text"`
	CreatedAt int64  `gorm:"autoCreateTime:milli"`
}

func init() {
	Models = append(Models, &OperationLog{})
}
