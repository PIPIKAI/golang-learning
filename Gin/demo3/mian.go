package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Articel struct {
	Title     string `json:"title"`
	Content   string `json:"content"`
	CreatTime int64  `json:"creatTime"`
}

func Unix2Time(ux int64) string {
	t := time.Unix(ux, 0)
	return t.Format("2014-01-02 15:04:05")
}
func main() {
	a := &Articel{
		Title:     "我是一个标题",
		Content:   "我是一个标题",
		CreatTime: time.Now().Unix(),
	}
	r := gin.Default()
	r.SetFuncMap(template.FuncMap{
		"Unix2Time": Unix2Time,
	})
	// 加载静态文件
	r.Static("/static", "./static")
	// 要成功使用html 的模板的话需要使用这句
	r.LoadHTMLGlob("templates/**/*")
	r.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "default/index.html", gin.H{
			"title":   "默认首页",
			"score":   89.9,
			"hobbys":  []string{"吃饭", "睡觉", "敲代码"},
			"articel": a,
		})
	})
	r.GET("/admin", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "admin/index.html", gin.H{
			"title": "admin首页",
		})
	})
	r.GET("/news", func(ctx *gin.Context) {

		ctx.HTML(http.StatusOK, "default/news.html", gin.H{
			"title": "新闻标题",
			"news":  a,
		})
	})
	fmt.Println("http://localhost:8080")
	r.Run()
}
