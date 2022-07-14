package admin

import (
	"github.com/gin-gonic/gin"
)

type NewsController struct{}

func (c NewsController) Index(ctx *gin.Context) {
	ctx.String(200, "admin--news")
}
