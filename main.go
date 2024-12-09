package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/testing", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"testing": "ok",
		})
	})

	r.Run(":5000")
}
