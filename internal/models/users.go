package model

import (
	"time"
)

// 用户表
type Users struct {
	UserID     uint      `gorm:"column:user_id;primary_key;AUTO_INCREMENT" json:"user_id"`               // 自增主键
	Nickname   string    `gorm:"column:nickname;NOT NULL" json:"nickname"`                               // 昵称
	AvatarUrl  string    `gorm:"column:avatar_url;NOT NULL" json:"avatar_url"`                           // 头像链接
	Openid     string    `gorm:"column:openid;NOT NULL" json:"openid"`                                   // 微信用户唯一标识
	Token      string    `gorm:"column:token;NOT NULL" json:"token"`                                     // 令牌
	UserStatus uint      `gorm:"column:user_status;default:1;NOT NULL" json:"user_status"`               // 用户状态 1:正常 2:禁止使用
	CreatedAt  time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP;NOT NULL" json:"created_at"` // 创建时间
	UpdatedAt  time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP;NOT NULL" json:"updated_at"` // 更新时间
}
