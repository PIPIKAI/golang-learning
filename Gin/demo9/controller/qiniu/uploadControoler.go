package qiniu

import (
	"bytes"
	"context"
	"demo9/models"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/qiniu/go-sdk/v7/storage"
)

type UploadControoler struct {
	BaseController
}

func (con UploadControoler) LocalUpload(ctx *gin.Context) {
	file, err := ctx.FormFile("file")
	var dst string
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	} else {
		extName := path.Ext(file.Filename)

		day := models.GetDay()

		tdir := path.Join("./static/upload", day)
		if err := os.MkdirAll(tdir, 0666); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "创建文件夹失败",
			})
			return
		} else {
			filename := strconv.FormatInt(models.GetUnix(), 10)
			dst = path.Join(tdir, filename+extName)
			fmt.Println(dst)
			fileContent, _ := file.Open()
			byteContainer, _ := ioutil.ReadAll(fileContent) // why the long names though?
			datalen := len(byteContainer)
			fmt.Printf("size:%d", datalen)
			// 保存到本地
			// ctx.SaveUploadedFile(file, dst)

			ret := storage.PutRet{}
			upToken, formUploader := con.Init()
			key := filename + extName
			err := formUploader.Put(context.Background(), &ret, upToken, key, bytes.NewReader(byteContainer), int64(datalen), nil)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(ret)
			sout := gin.H{
				"message":   "上传成功",
				"path":      dst,
				"qiniu_url": con.PublicAccessURL(key),
			}
			ctx.JSON(http.StatusBadRequest, sout)
		}

	}

}
