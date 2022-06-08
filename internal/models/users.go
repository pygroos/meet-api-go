package models

import (
	"time"
)

type Users struct {
	UserID     uint      `gorm:"column:user_id;primary_key;AUTO_INCREMENT"`            // 自增主键
	Openid     string    `gorm:"column:openid;NOT NULL"`                               // 微信用户唯一标识
	Phone      string    `gorm:"column:phone;NOT NULL"`                                // 手机号
	Nickname   string    `gorm:"column:nickname;NOT NULL"`                             // 昵称
	AvatarUrl  string    `gorm:"column:avatar_url;NOT NULL"`                           // 头像链接
	Model      string    `gorm:"column:model;NOT NULL"`                                // 车型
	Token      string    `gorm:"column:token;NOT NULL"`                                // 令牌
	UserStatus uint      `gorm:"column:user_status;default:1;NOT NULL"`                // 用户状态 1:正常 2:禁止使用
	CreatedAt  time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP;NOT NULL"` // 创建时间
	UpdatedAt  time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP;NOT NULL"` // 更新时间
}

func (m *Users) TableName() string {
	return "users"
}