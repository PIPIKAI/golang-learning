package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
			"data":    132,
		})
	})
	r.POST("/newpost", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"title":  "title",
			"contex": 132,
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080

}
