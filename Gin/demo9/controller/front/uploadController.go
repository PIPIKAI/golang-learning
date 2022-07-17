package front

import (
	"demo9/models"
	"fmt"
	"net/http"
	"os"
	"path"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UploadController struct{}

func (con UploadController) Init(ctx *gin.Context) {
	ctx.HTML(200, "default/upload.html", gin.H{
		"title": "上传图片",
	})
}

func (con UploadController) MulInit(ctx *gin.Context) {
	ctx.HTML(200, "default/mulupload.html", gin.H{
		"title": "多文件上传图片",
	})
}
func (con UploadController) SignelUpload(ctx *gin.Context) {
	username := ctx.PostForm("username")
	// 单文件上传
	file, err := ctx.FormFile("face")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	} else {
		dst := path.Join("./static/upload", file.Filename)
		fmt.Println(dst)
		ctx.SaveUploadedFile(file, dst)
		ctx.JSON(200, gin.H{
			"username": username,
			"filename": fmt.Sprintf("%v uploaded! ", file.Filename),
		})
	}

}

func (con UploadController) MulUpload(ctx *gin.Context) {
	username := ctx.PostForm("username")
	// 多文件上传
	form, err := ctx.MultipartForm()
	files := form.File["face[]"]
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	} else {
		allowExtMap := map[string]bool{".jpg": true, ".png": true, ".gif": true, ".jpeg": true}
		for _, file := range files {
			extName := path.Ext(file.Filename)
			if allowExtMap[extName] {

				day := models.GetDay()

				tdir := path.Join("./static/upload", day)
				if err := os.MkdirAll(tdir, 0666); err != nil {
					ctx.JSON(http.StatusBadRequest, gin.H{
						"message": "创建文件夹失败",
					})
					return
				} else {
					filename := strconv.FormatInt(models.GetUnix(), 10)
					dst := path.Join(tdir, filename+extName)
					fmt.Println(dst)
					ctx.SaveUploadedFile(file, dst)
				}

			} else {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"message": "文件上传失败！不支持此类型的文件",
				})
				return
			}

		}

		ctx.JSON(200, gin.H{
			"message":  "文件上传成功",
			"username": username,
		})
	}

}
