package models

import (
	"time"
)

type Records struct {
	RecordID     uint      `gorm:"column:record_id;primary_key;AUTO_INCREMENT"`          // 自增主键
	ActivityID   uint      `gorm:"column:activity_id;NOT NULL"`                          // 活动id
	UserID       uint      `gorm:"column:user_id;NOT NULL"`                              // 用户id
	RecordStatus uint      `gorm:"column:record_status;default:1;NOT NULL"`              // 状态 1:正常 2:取消
	CreatedAt    time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP;NOT NULL"` // 创建时间
	UpdatedAt    time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP;NOT NULL"` // 更新时间
}

func (m *Records) TableName() string {
	return "records"
}