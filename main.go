package main

import "fmt"
import "github.com/gin-gonic/gin"

func main() {
	fmt.Println("vim-go")
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
			"code":    "1",
		})
	})
	r.Run()
}
