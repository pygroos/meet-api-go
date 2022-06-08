package models

import (
	"time"
)

type Activities struct {
	ActivityID         uint      `gorm:"column:activity_id;primary_key;AUTO_INCREMENT"`          // 自增主键
	Subject            string    `gorm:"column:subject;NOT NULL"`                                // 标题
	Content            string    `gorm:"column:content;NOT NULL"`                                // 内容
	Province           string    `gorm:"column:province;NOT NULL"`                               // 省
	City               string    `gorm:"column:city;NOT NULL"`                                   // 市
	StartTime          time.Time `gorm:"column:start_time;default:1000-01-01 00:00:00;NOT NULL"` // 开始时间
	EndTime            time.Time `gorm:"column:end_time;default:1000-01-01 00:00:00;NOT NULL"`   // 结束时间
	WeekFlag           uint      `gorm:"column:week_flag;default:0;NOT NULL"`                    // 周 1~7
	MeetPlace          string    `gorm:"column:meet_place;NOT NULL"`                             // 集合地点
	MeetPlaceLongitude string    `gorm:"column:meet_place_longitude;default:0.000000;NOT NULL"`  // 集合地点经度
	MeetPlaceLatitude  string    `gorm:"column:meet_place_latitude;default:0.000000;NOT NULL"`   // 集合地点纬度
	PublisherID        int       `gorm:"column:publisher_id;default:0;NOT NULL"`                 // 发布人id
	ImagesUrl          string    `gorm:"column:images_url;NOT NULL"`                             // 图片链接
	PeopleNumberLimit  uint      `gorm:"column:people_number_limit;default:20;NOT NULL"`         // 报名人数限制
	ActivityStatus     uint      `gorm:"column:activity_status;default:1;NOT NULL"`              // 状态 1:正常 2:结束 3:取消
	CreatedAt          time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP;NOT NULL"`   // 创建时间
	UpdatedAt          time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP;NOT NULL"`   // 更新时间
}

func (m *Activities) TableName() string {
	return "activities"
}