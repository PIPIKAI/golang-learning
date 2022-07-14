package admin

import (
	"github.com/gin-gonic/gin"
)

type UserController struct {
	BaseController
}

func (c UserController) Index(ctx *gin.Context) {

	ctx.String(200, "admin--user")
	c.Success(ctx)
}
