package route

import (
	"github.com/gin-gonic/gin"
	"meet-api/internal/api"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	var user api.UserController
	var meet api.MeetController
	var location api.LocationController
	apiV1 := r.Group("/api/v1")
	{
		// 登录
		apiV1.POST("/user/login", user.Login)
		// 用户信息
		apiV1.GET("/user/info", user.Info)
		// 遇见记录
		apiV1.GET("meet/list", meet.List)
		// 遇见详情
		apiV1.GET("meet/info", meet.Info)
		// 位置打点
		apiV1.POST("location/dot", location.Dot)
	}

	return r
}
