package routers

import (
	"demo7/controller/front"
	"demo7/middleware"

	"github.com/gin-gonic/gin"
)

func DefaultRoutersInit(r *gin.Engine) {
	Router := r.Group("/")
	{
		Router.GET("/", middleware.RecoderTiem, front.DefaultController{}.Init)
	}
}
