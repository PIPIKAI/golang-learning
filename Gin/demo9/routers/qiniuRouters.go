package routers

import (
	"demo9/controller/qiniu"
	"demo9/middleware"

	"github.com/gin-gonic/gin"
)

func QiniuRoutersInit(r *gin.Engine) {
	Router := r.Group("/qiniu")
	{
		Router.GET("/", qiniu.FrontController{}.Index)
		Router.POST("/upload", middleware.RecoderTiem, qiniu.UploadControoler{}.LocalUpload)
	}
}
