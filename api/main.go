package main

import (
	"github.com/gin-gonic/gin"
	"github.com/torbenconto/plutus/quote"
	"net/http"
)

func setupRouter() *gin.Engine {
	// Initialize gin router
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/quote/:ticker", func(c *gin.Context) {
		// Get ticker from url param
		ticker := c.Param("ticker")
		// Create new quote instance
		stock, err := quote.NewQuote(ticker)
		// Check for errors, return 404 if not found or 200 along with quote data if found
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": err,
			})
		} else {
			c.JSON(http.StatusOK, stock)
		}

	})

	return r
}

func main() {
	r := setupRouter()

	// Listen and serve on 8080
	err := r.Run(":8080")
	if err != nil {
		return
	}
}
