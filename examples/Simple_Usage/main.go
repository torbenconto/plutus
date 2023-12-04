package main

import (
	"fmt"

	"github.com/torbenconto/plutus"
)

func main() {
	// Create a stock object using a ticker, the stock will be auto populated with data when created, no need to call any other methods
	// Returns an error and a plutus.Stock struct, error details what went wrong with fetching data
	//                  no need to capitalize ticker
	//                              |
	//                              v
	stock, err := plutus.NewStock("amd")
	if err != nil {
		fmt.Printf("An error occured: %e", err)
	}

	// The resulting stock object has many different fields of data filled in from yahoo finance
	fmt.Printf("Stock Information:\n")
	fmt.Printf("Ticker: %s\n", stock.Ticker)
	fmt.Printf("Price: %.2f\n", stock.Price)
	fmt.Printf("Change: %.2f\n", stock.Change)
	fmt.Printf("Prev Close: %.2f\n", stock.PrevClose)
	fmt.Printf("Open Price: %.2f\n", stock.OpenPrice)
	fmt.Printf("Bid Price: %.2f\n", stock.BidPrice)
	fmt.Printf("Ask Price: %.2f\n", stock.AskPrice)
	fmt.Printf("Day Range: %s\n", stock.DayRange)
	fmt.Printf("52-Week Range: %s\n", stock.FiftyTwoWeekRange)
	fmt.Printf("Volume: %d\n", stock.Volume)
	fmt.Printf("Average Volume: %d\n", stock.AvgVolume)
	fmt.Printf("Market Cap: %s\n", stock.MarketCap)
	fmt.Printf("Beta: %.2f\n", stock.Beta)
	fmt.Printf("PE Ratio: %.2f\n", stock.PE)
	fmt.Printf("EPS: %.2f\n", stock.EPS)
	fmt.Printf("Earnings Date: %s\n", stock.EarningsDate)
	fmt.Printf("Forward Dividend & Yield: %s\n", stock.ForwardDividendAndYield)
	fmt.Printf("Ex-Dividend Date: %s\n", stock.ExDividendDate)
	fmt.Printf("1-Year Target Estimate: %.2f\n", stock.OneYearTargetEst)

	// If you want a constant stream of data on the stock you can use the Stream method on the stock
	// An example of this method is contained here
	// https://github.com/torbenconto/plutus/blob/master/examples/Stock_Data_Stream/main.go
}
