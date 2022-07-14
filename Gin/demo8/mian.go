package main

import (
	"demo8/routers"
	"fmt"
	"html/template"
	"time"

	"github.com/gin-gonic/gin"
)

type Article struct {
	Title   string `json:"title" xml:"title"`
	Content string `json:"content" xml:"content"`
}
type UserInfo struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
	Age      int64  `json:"age" form:"age"`
}

func Unix2Time(ux int64) string {
	t := time.Unix(ux, 0)
	return t.Format("2014-01-02 15:04:05")
}
func main() {
	r := gin.Default()
	r.SetFuncMap(template.FuncMap{
		"Unix2Time": Unix2Time,
	})
	// 加载静态文件
	r.Static("/static", "./static")
	// 要成功使用html 的模板的话需要使用这句
	r.LoadHTMLGlob("templates/**/*")

	// 使用Get 传值
	routers.DefaultRoutersInit(r)
	routers.AdminRoutersInit(r)
	routers.ApiRoutersInit(r)

	fmt.Println("http://localhost:8080")
	r.Run()
}
