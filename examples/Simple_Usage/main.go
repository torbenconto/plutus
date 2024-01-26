package main

import (
	"fmt"
	"github.com/torbenconto/plutus/historical"
	"github.com/torbenconto/plutus/interval"
	"github.com/torbenconto/plutus/range"
)

func main() {
	// Create a quote object using a ticker, the quote will be auto populated with data when created, no need to call any other methods
	// Returns an error and a plutus.Stock struct, error details what went wrong with fetching data
	//                  no need to capitalize ticker
	//                              |
	//                              v
	stock, err := historical.NewHistorical("amd", _range.FiveDay, interval.ThirtyMinutes)
	if err != nil {
		fmt.Printf("An error occured: %e", err)
	}

	fmt.Println(stock)
}

// If you want a constant stream of data on the quote you can use the Stream method on the quote
// An example of this method is contained here
// https://github.com/torbenconto/plutus/blob/master/examples/Stock_Data_Stream/main.go
