package api

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"meet-api/internal/pkg/support"
	"meet-api/internal/services"
)

func GetActivityList(c *gin.Context) {
	r := support.Gin{C: c}
	city := c.Query("city")
	week, _ := strconv.Atoi(c.Query("week"))

	activityService := services.ActivityService{}

	list, total := activityService.GetActivityList(city, week)
	r.Ok(map[string]interface{}{
		"list":  list,
		"total": total,
	})
}
