package routers

import (
	"demo8/controller/admin"
	"demo8/middleware"

	"github.com/gin-gonic/gin"
)

func AdminRoutersInit(r *gin.Engine) {
	Router := r.Group("/admin")
	{
		Router.GET("/user", middleware.RecoderTiem, admin.UserController{}.Index)
		Router.GET("/news", admin.NewsController{}.Index)
	}
}
