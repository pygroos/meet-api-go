package models

import (
	"time"
)

type Activities struct {
	ActivityID         uint      `gorm:"column:activity_id;primary_key;AUTO_INCREMENT" json:"activity_id"`                  // 自增主键
	Subject            string    `gorm:"column:subject;NOT NULL" json:"subject"`                                            // 标题
	Content            string    `gorm:"column:content;NOT NULL" json:"content"`                                            // 内容
	Province           string    `gorm:"column:province;NOT NULL" json:"province"`                                          // 省
	City               string    `gorm:"column:city;NOT NULL" json:"city"`                                                  // 市
	StartTime          time.Time `gorm:"column:start_time;default:1000-01-01 00:00:00;NOT NULL" json:"start_time"`          // 开始时间
	EndTime            time.Time `gorm:"column:end_time;default:1000-01-01 00:00:00;NOT NULL" json:"end_time"`              // 结束时间
	WeekFlag           uint      `gorm:"column:week_flag;default:0;NOT NULL" json:"week_flag"`                              // 周 1~7
	MeetPlace          string    `gorm:"column:meet_place;NOT NULL" json:"meet_place"`                                      // 集合地点
	MeetPlaceLongitude string    `gorm:"column:meet_place_longitude;default:0.000000;NOT NULL" json:"meet_place_longitude"` // 集合地点经度
	MeetPlaceLatitude  string    `gorm:"column:meet_place_latitude;default:0.000000;NOT NULL" json:"meet_place_latitude"`   // 集合地点纬度
	PublisherID        int       `gorm:"column:publisher_id;default:0;NOT NULL" json:"publisher_id"`                        // 发布人id
	ImagesUrl          string    `gorm:"column:images_url;NOT NULL" json:"images_url"`                                      // 图片链接
	PeopleNumberLimit  uint      `gorm:"column:people_number_limit;default:20;NOT NULL" json:"people_number_limit"`         // 报名人数限制
	ActivityStatus     uint      `gorm:"column:activity_status;default:1;NOT NULL" json:"activity_status"`                  // 状态 1:正常 2:结束 3:取消
	CreatedAt          time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP;NOT NULL" json:"created_at"`            // 创建时间
	UpdatedAt          time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP;NOT NULL" json:"updated_at"`            // 更新时间
}
