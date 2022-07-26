package repositories

import (
	. "meet-api/internal/models"
	. "meet-api/internal/pkg/database"
)

const (
	ACTIVITY_NORMAL = 1
	ACTIVITY_END    = 2
	ACTIVITY_CANCEL = 3
)

func GetListByConditions(city string, week int) ([]Activities, int) {
	var list []Activities
	var total int

	tx := DbHandle
	if city != "" {
		tx = tx.Where("city = (?)", []string{city})
	}
	if week != 0 {
		tx = tx.Where("week_flag = (?)", []int{week})
	}

	tx.Find(&list)
	tx.Find(&list).Count(&total)

	return list, total
}
