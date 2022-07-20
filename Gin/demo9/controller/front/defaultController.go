package front

import "github.com/gin-gonic/gin"

type DefaultController struct{}

func (con DefaultController) Init(ctx *gin.Context) {
	// 设置cookie
	ctx.SetCookie("name", "这是一个cookie", 3600, "/", "localhost", false, true)
	ctx.HTML(200, "default/index.html", gin.H{
		"title": "我是一个默认首页",
	})
}
