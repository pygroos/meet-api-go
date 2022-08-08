package model

import (
	"time"
)

// 遇见记录表
type Meets struct {
	ID                  uint      `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`                                      // 自增主键
	MeetDate            string    `gorm:"column:meet_date;NOT NULL" json:"meet_date"`                                          // 遇见日期
	MeetTime            uint      `gorm:"column:meet_time;default:0;NOT NULL" json:"meet_time"`                                // 遇见时间戳
	FirstUserID         uint      `gorm:"column:first_user_id;default:0;NOT NULL" json:"first_user_id"`                        // 第一位用户ID
	FirstUserLongitude  string    `gorm:"column:first_user_longitude;default:0.000000;NOT NULL" json:"first_user_longitude"`   // 第一位用户经度
	FirstUserLatitude   string    `gorm:"column:first_user_latitude;default:0.000000;NOT NULL" json:"first_user_latitude"`     // 第一位用户纬度
	SecondUserID        uint      `gorm:"column:second_user_id;default:0;NOT NULL" json:"second_user_id"`                      // 第二位用户ID
	SecondUserLongitude string    `gorm:"column:second_user_longitude;default:0.000000;NOT NULL" json:"second_user_longitude"` // 第二位用户经度
	SecondUserLatitude  string    `gorm:"column:second_user_latitude;default:0.000000;NOT NULL" json:"second_user_latitude"`   // 第二位用户纬度
	IsViewed            uint      `gorm:"column:is_viewed;default:0;NOT NULL" json:"is_viewed"`                                // 是否查看 0:否 1:是
	CreatedAt           time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP;NOT NULL" json:"created_at"`              // 创建时间
	UpdatedAt           time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP;NOT NULL" json:"updated_at"`              // 更新时间
}
