package qiniu

import (
	"github.com/gin-gonic/gin"
)

type FrontController struct {
}

func (con FrontController) Index(ctx *gin.Context) {
	ctx.HTML(200, "qiniu/index.html", gin.H{
		"title": "七牛操作demo",
	})
}
