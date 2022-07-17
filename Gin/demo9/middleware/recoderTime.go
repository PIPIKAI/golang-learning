package middleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func RecoderTiem(ctx *gin.Context) {
	start := time.Now()
	// 当在中间件或 handler 中启动新的 goroutine 时，不能使用原始的上下文（c *gin.Context），必须使用其只读副本（c.Copy()）
	cpC := ctx.Copy()
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("cpC.Request.URL.Path : ", cpC.Request.URL.Path)
	}()
	ctx.Next()
	fmt.Println("MiddleWare using ---:", time.Now().Sub(start).Milliseconds(), "ms")
}
