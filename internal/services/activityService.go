package services

import (
	. "meet-api/internal/models"
	"meet-api/internal/repositories"
)

type ActivityService struct{}

func (a *ActivityService) GetActivityList(city string, week int) ([]Activities, int) {
	var list []Activities

	// TODO 组装查询条件放在service层
	list, total := repositories.GetListByConditions(city, week)

	return list, total
}
