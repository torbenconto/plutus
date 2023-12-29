package main

import (
	"github.com/gin-gonic/gin"
	"github.com/torbenconto/plutus"
)

func setupRouter() *gin.Engine {
	// Initialize gin router
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/stock/:ticker", func(c *gin.Context) {
		// Get provider from url param

		var activeProvider plutus.StockDataProvider
		providerQuery := c.Query("provider")
		// None specified, use default
		if providerQuery == "" {
			activeProvider = plutus.YahooFinanceProvider
		}

		ticker := c.Param("ticker")
		stock, err := plutus.NewStock(ticker, activeProvider)
		if err != nil {
			c.JSON(404, gin.H{
				"error": err,
			})
		} else {
			c.JSON(200, stock)
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
