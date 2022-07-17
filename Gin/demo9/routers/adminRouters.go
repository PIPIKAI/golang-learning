package routers

import (
	"demo9/controller/admin"
	"demo9/middleware"

	"github.com/gin-gonic/gin"
)

func AdminRoutersInit(r *gin.Engine) {
	Router := r.Group("/admin")
	{
		Router.GET("/user", middleware.RecoderTiem, admin.UserController{}.Index)
		Router.GET("/news", admin.NewsController{}.Index)
	}
}
