package main

import (
	"demo9/routers"
	"fmt"
	"html/template"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
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
	// session 的使用
	// 加密密钥
	encodePassword := "secret123"
	// 保存在本地的session
	// store := cookie.NewStore([]byte(encodePassword))
	// 保存在redis中的session
	// 初始化基于 redis 的存储引擎
	// 参数说明：
	//第 1 个参数 - redis 最大的空闲连接数
	//第 2 个参数 - 数通信协议 tcp 或者 udp
	//第 3 个参数 - redis 地址, 格式，host:port
	//第 4 个参数 - redis 密码
	//第 5 个参数 - session 加密密钥
	store, _ := redis.NewStore(10, "tcp", "localhost:6379", "", []byte(encodePassword))
	// 使用session
	r.Use(sessions.Sessions("mysession", store))
	// 加载静态文件
	r.Static("/static", "./static")
	// 要成功使用html 的模板的话需要使用这句
	r.LoadHTMLGlob("templates/**/*")

	// 设置cookie
	// 使用Get 传值
	routers.DefaultRoutersInit(r)
	routers.AdminRoutersInit(r)
	routers.ApiRoutersInit(r)
	routers.QiniuRoutersInit(r)
	fmt.Println("http://localhost:8080")
	r.Run()
}
