package api

import (
	"github.com/gin-gonic/gin"
	"meet-api/internal/code"
	"meet-api/internal/pkg/support"
	"meet-api/internal/services"
	"strconv"
)

type UserController struct {
}

func (u *UserController) Login(ctx *gin.Context) {
	wxCode := ctx.Param("code")
	userSrv := services.UserService{}
	res := userSrv.Login(wxCode)
	r := support.Gin{C: ctx}
	if res {
		r.Ok(map[string]string{})
	} else {
		r.Fail(code.LOGIN_FAILED)
	}
}

func (u *UserController) Info(ctx *gin.Context) {
	userId, _ := strconv.Atoi(ctx.Query("user_id"))
	userSrv := services.UserService{}
	r := support.Gin{C: ctx}
	r.Ok(userSrv.Info(userId))
}
