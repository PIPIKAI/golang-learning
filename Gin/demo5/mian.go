package main

import (
	"encoding/xml"
	"fmt"
	"html/template"
	"net/http"
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
	r.GET("/user", func(ctx *gin.Context) {
		uid := ctx.Query("uid")
		page := ctx.DefaultQuery("page", "0")
		ctx.JSON(http.StatusOK, gin.H{
			"uid":  uid,
			"page": page,
		})
	})
	r.GET("/regiser", func(ctx *gin.Context) {
		ctx.HTML(200, "default/user.html", gin.H{})
	})
	// 获取表单的数据
	r.POST("/Adduer", func(ctx *gin.Context) {
		username := ctx.PostForm("username")
		password := ctx.PostForm("password")
		age := ctx.DefaultPostForm("age", "20")
		ctx.JSON(http.StatusOK, gin.H{
			"username": username,
			"password": password,
			"age":      age,
		})
	})
	// 使用结构体来获得表单的数据
	r.GET("/GetUser", func(ctx *gin.Context) {
		user := &UserInfo{}

		if err := ctx.ShouldBind(&user); err == nil {
			ctx.JSON(200, user)
		} else {
			ctx.JSON(200, gin.H{
				"err": err,
			})
		}
	})
	// 获取表单的数据
	r.POST("/Adduer2", func(ctx *gin.Context) {
		user := &UserInfo{}

		if err := ctx.ShouldBind(&user); err == nil {
			ctx.JSON(200, user)
		} else {
			ctx.JSON(200, gin.H{
				"err": err,
			})
		}
	})
	// 获取 Post Xml 数据
	r.POST("/xml", func(ctx *gin.Context) {
		b, _ := ctx.GetRawData()

		articel := &Article{}
		if err := xml.Unmarshal(b, articel); err == nil {
			ctx.JSON(200, articel)
		} else {
			ctx.JSON(http.StatusBadRequest, err.Error())
		}
	})
	// 动态路由
	r.GET("/user/:uid", func(ctx *gin.Context) {
		uid := ctx.Param("uid")

		ctx.String(200, "uid=%s", uid)
	})
	fmt.Println("http://localhost:8080")
	r.Run()
}
