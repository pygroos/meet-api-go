package support

import (
	"github.com/gin-gonic/gin"
	"meet-api/internal/code"
	"net/http"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func (g *Gin) Ok(data interface{}) {
	g.C.JSON(http.StatusOK, Response{
		Code: code.SUCCESS,
		Msg:  code.GetMsg(code.SUCCESS),
		Data: data,
	})
	return
}

func (g *Gin) Fail(errCode int) {
	g.C.JSON(http.StatusOK, Response{
		Code: errCode,
		Msg:  code.GetMsg(errCode),
		Data: "",
	})
	return
}
