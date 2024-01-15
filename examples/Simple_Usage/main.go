package main

import (
	"fmt"
	quote "github.com/torbenconto/plutus/quote"
)

func main() {
	// Create a quote object using a ticker, the quote will be auto populated with data when created, no need to call any other methods
	// Returns an error and a plutus.Stock struct, error details what went wrong with fetching data
	//                  no need to capitalize ticker
	//                              |
	//                              v
	stock, err := quote.NewQuote("amd")
	if err != nil {
		fmt.Printf("An error occured: %e", err)
	}

	// The resulting quote object has many different fields of data filled in from yahoo finance
	fmt.Printf("Ticker: %s\n", stock.Ticker)
	fmt.Printf("Price: %.2f\n", stock.Price)
	fmt.Printf("ChangePrice: %.2f\n", stock.ChangePrice)
	fmt.Printf("ChangePercent: %.2f%%\n", stock.ChangePercent)
	fmt.Printf("PrevClose: %.2f\n", stock.PrevClose)
	fmt.Printf("OpenPrice: %.2f\n", stock.OpenPrice)
	fmt.Printf("BidPrice: %s\n", stock.BidPrice)
	fmt.Printf("AskPrice: %s\n", stock.AskPrice)
	fmt.Printf("DayHigh: %.2f\n", stock.DayHigh)
	fmt.Printf("DayLow: %.2f\n", stock.DayLow)
	fmt.Printf("FiftyTwoWeekLow: %.2f\n", stock.FiftyTwoWeekLow)
	fmt.Printf("FiftyTwoWeekHigh: %.2f\n", stock.FiftyTwoWeekHigh)
	fmt.Printf("Volume: %d\n", stock.Volume)
	fmt.Printf("AvgVolume: %d\n", stock.AvgVolume)
	fmt.Printf("MarketCap: %s\n", stock.MarketCap)
	fmt.Printf("Beta: %.2f\n", stock.Beta)
	fmt.Printf("PE: %.2f\n", stock.PE)
	fmt.Printf("EPS: %.2f\n", stock.EPS)
	fmt.Printf("EarningsDate: %s\n", stock.EarningsDate)
	fmt.Printf("ForwardDividendAndYield: %s\n", stock.ForwardDividendAndYield)
	fmt.Printf("DividendDate: %s\n", stock.DividendDate)
	fmt.Printf("ExDividendDate: %s\n", stock.ExDividendDate)
	fmt.Printf("OneYearTargetEst: %.2f\n", stock.OneYearTargetEst)

	// If you want a constant stream of data on the quote you can use the Stream method on the quote
	// An example of this method is contained here
	// https://github.com/torbenconto/plutus/blob/master/examples/Stock_Data_Stream/main.go
}
