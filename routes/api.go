package route

import (
	"meet-api/internal/api"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	apiV1 := r.Group("/api/v1/activity")
	{
		// 活动列表
		apiV1.GET("/list", api.GetActivityList)
		// 活动详情
		// apiV1.GET("/detail", api.GetActivityDetail)
	}

	return r
}
