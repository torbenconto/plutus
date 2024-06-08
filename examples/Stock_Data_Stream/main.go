package main

import (
	"fmt"
	"github.com/torbenconto/plutus/stock"
	"time"
)

func main() {
	// Create new Stock or Historical  object
	data, err := stock.NewQuote("GOOG")
	if err != nil {
		fmt.Println("Error fetching data for stock")
	}

	// Set delay in Milliseconds (1000 = 1 second)
	delay := time.Second

	// Call stream func using Stock object and a given delay
	stream := data.Stream(delay)

	// Get updated data and print out most recent stock price. Runs infinently and returns the newest avalible stock data in the form of a plutus.stock.Quote struct
	for {
		streamData := <-stream
		fmt.Println(streamData.RegularMarketPrice, streamData.RegularMarketChangePercent)
	}

}
