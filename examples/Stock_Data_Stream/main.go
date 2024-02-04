package main

import (
	"fmt"
	"github.com/torbenconto/plutus/quote"
	"time"
)

func main() {
	// Create new Stock or Historical  object
	stock, err := quote.NewQuote("GOOG")
	if err != nil {
		fmt.Println("Error fetching data for quote")
	}

	// Set delay in Milliseconds (1000 = 1 second)
	delay := time.Second

	// Call stream func using Stock object and a given delay
	stream := stock.Stream(delay)

	// Get updated data and print out most recent quote price. Runs infinently and returns the newest avalible quote data in the form of a plutus.Stock struct
	for {
		data := <-stream
		fmt.Println(data.RegularMarketPrice, data.RegularMarketChangePercent)
	}

}
