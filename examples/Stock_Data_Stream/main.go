package main

import (
	"fmt"
	"github.com/torbenconto/plutus/quote"
)

func main() {
	// Create new Stock object
	stock, err := quote.NewQuote("AMD")
	if err != nil {
		fmt.Println("Error fetching data for quote")
	}

	// Set delay in Milliseconds
	delayInMS := 1000

	// Call stream func using Stock object and a given delay
	stream := stock.Stream(delayInMS)

	// Get updated data and print out most recent quote price. Runs infinently and returns the newest avalible quote data in the form of a plutus.Stock struct
	for {
		data := <-stream
		fmt.Println(data.Price, data.ChangePercent)
	}

}
