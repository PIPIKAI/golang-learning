package routers

import (
	"demo9/controller/front"
	"demo9/middleware"

	"github.com/gin-gonic/gin"
)

func DefaultRoutersInit(r *gin.Engine) {
	Router := r.Group("/")
	{
		Router.GET("/", middleware.RecoderTiem, front.DefaultController{}.Init)
		Router.POST("/upload", middleware.RecoderTiem, front.UploadController{}.SignelUpload)
		Router.GET("/edit", middleware.RecoderTiem, front.UploadController{}.Init)

		Router.GET("/pic", front.UploadController{}.MulInit)
		Router.POST("/mulupload", front.UploadController{}.MulUpload)
	}
}
