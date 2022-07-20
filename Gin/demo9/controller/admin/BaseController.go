package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BaseController struct {
}

func (c BaseController) Success(ctx *gin.Context) {
	// 获取cookie
	v, err := ctx.Cookie("name")
	if err != nil {
		ctx.String(http.StatusBadRequest, "失败 , Error %v", err)
	}
	ctx.String(http.StatusOK, "成功 , cookie name : name \n cookie value : %v", v)
}
func (c BaseController) Error(ctx *gin.Context) {
	ctx.String(http.StatusOK, "失败")
}
