package front

import "github.com/gin-gonic/gin"

type DefaultController struct{}

func (con DefaultController) Init(ctx *gin.Context) {
	ctx.HTML(200, "default/index.html", gin.H{
		"title": "我是一个默认首页",
	})
}
