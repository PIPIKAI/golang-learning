package routers

import "github.com/gin-gonic/gin"

func ApiRoutersInit(r *gin.Engine) {
	Router := r.Group("/api")
	{
		Router.GET("/", func(ctx *gin.Context) {
			ctx.String(200, "api--")
		})
		Router.GET("/user", func(ctx *gin.Context) {
			ctx.String(200, "api--user")
		})
		Router.GET("/admin", func(ctx *gin.Context) {
			ctx.String(200, "api--admin")
		})
	}
}
