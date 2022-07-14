package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Articel struct {
	Title     string        `json:"title"`
	Content   string        `json:"content"`
	CreatTime time.Duration `json:"creatTime"`
}

func main() {
	a := Articel{
		Title:     "我是一个标题",
		Content:   "我是一个标题",
		CreatTime: time.Duration(time.Now().Unix()),
	}
	r := gin.Default()
	// 要成功使用html 的模板的话需要使用这句
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "the root")
	})
	r.GET("/json1", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"title":     "标题",
			"creattime": time.Now().Unix(),
		})
	})
	r.GET("/json2", func(ctx *gin.Context) {
		ctx.JSON(200, a)
	})
	// xxxxcallback({"creattime":1656741087,"title":"标题"});
	r.GET("/jsonp", func(ctx *gin.Context) {
		ctx.JSONP(200, gin.H{
			"title":     "标题",
			"creattime": time.Now().Unix(),
		})
	})
	//
	r.GET("/xml", func(ctx *gin.Context) {
		ctx.XML(http.StatusOK, a)
	})

	r.GET("/articel", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "articel.html", gin.H{
			"title": "我是标题",
		})
	})
	fmt.Println("http://localhost:8080")
	r.Run()
}
