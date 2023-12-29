package main

import (
	"github.com/gin-gonic/gin"
)

func setupRouter() {
	// Initialize gin router
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}

func main() {
}
