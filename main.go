package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {

	// 日志
	//gin.DisableConsoleColor()
	//f, _ := os.Create("gin.log")
	//gin.DefaultWriter = io.MultiWriter(f)

	log.Println("====日志===")

	r := gin.Default()
	gin.Logger()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
