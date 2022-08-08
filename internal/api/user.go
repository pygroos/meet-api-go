package api

import (
	"github.com/gin-gonic/gin"
	"meet-api/internal/pkg/support"
	"meet-api/internal/services"
	"strconv"
)

type UserController struct {
}

func (u *UserController) Login(ctx *gin.Context) {
	code := ctx.Param("code")
	userSrv := services.UserService{}
	res := userSrv.Login(code)
	r := support.Gin{C: ctx}
	r.Ok(map[string]interface{}{
		"res": res,
	})
}

func (u *UserController) Info(ctx *gin.Context) {
	userId, _ := strconv.Atoi(ctx.Query("user_id"))
	userSrv := services.UserService{}
	r := support.Gin{C: ctx}
	r.Ok(userSrv.Info(userId))
}
