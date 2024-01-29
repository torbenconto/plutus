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
		fmt.Printf("An error occured: %s\n", err)
	}

	// Print all values defined in the Quote struct
	fmt.Printf("Language: %s\n", stock.Language)
	fmt.Printf("Region: %s\n", stock.Region)
	fmt.Printf("QuoteType: %s\n", stock.QuoteType)
	fmt.Printf("TypeDisp: %s\n", stock.TypeDisp)
	fmt.Printf("QuoteSourceName: %s\n", stock.QuoteSourceName)
	fmt.Printf("Triggerable: %v\n", stock.Triggerable)
	fmt.Printf("CustomPriceAlertConfidence: %s\n", stock.CustomPriceAlertConfidence)
	fmt.Printf("Currency: %s\n", stock.Currency)
	fmt.Printf("MarketState: %s\n", stock.MarketState)
	fmt.Printf("RegularMarketChangePercent: %.2f\n", stock.RegularMarketChangePercent)
	fmt.Printf("RegularMarketPrice: %.2f\n", stock.RegularMarketPrice)
	fmt.Printf("Exchange: %s\n", stock.Exchange)
	fmt.Printf("ShortName: %s\n", stock.ShortName)
	fmt.Printf("LongName: %s\n", stock.LongName)
	fmt.Printf("MessageBoardID: %s\n", stock.MessageBoardID)
	fmt.Printf("ExchangeTimezoneName: %s\n", stock.ExchangeTimezoneName)
	fmt.Printf("ExchangeTimezoneShortName: %s\n", stock.ExchangeTimezoneShortName)
	fmt.Printf("GmtOffSetMilliseconds: %d\n", stock.GmtOffSetMilliseconds)
	fmt.Printf("Market: %s\n", stock.Market)
	fmt.Printf("EsgPopulated: %v\n", stock.EsgPopulated)
	fmt.Printf("FirstTradeDateMilliseconds: %d\n", stock.FirstTradeDateMilliseconds)
	fmt.Printf("PriceHint: %d\n", stock.PriceHint)
	fmt.Printf("PostMarketChangePercent: %.2f\n", stock.PostMarketChangePercent)
	fmt.Printf("PostMarketTime: %d\n", stock.PostMarketTime)
	fmt.Printf("PostMarketPrice: %.2f\n", stock.PostMarketPrice)
	fmt.Printf("PostMarketChange: %.2f\n", stock.PostMarketChange)
	fmt.Printf("RegularMarketChange: %.2f\n", stock.RegularMarketChange)
	fmt.Printf("RegularMarketTime: %d\n", stock.RegularMarketTime)
	fmt.Printf("RegularMarketDayHigh: %.2f\n", stock.RegularMarketDayHigh)
	fmt.Printf("RegularMarketDayRange: %s\n", stock.RegularMarketDayRange)
	fmt.Printf("RegularMarketDayLow: %.2f\n", stock.RegularMarketDayLow)
	fmt.Printf("RegularMarketVolume: %d\n", stock.RegularMarketVolume)
	fmt.Printf("RegularMarketPreviousClose: %.2f\n", stock.RegularMarketPreviousClose)
	fmt.Printf("Bid: %.2f\n", stock.Bid)
	fmt.Printf("Ask: %.2f\n", stock.Ask)
	fmt.Printf("BidSize: %d\n", stock.BidSize)
	fmt.Printf("AskSize: %d\n", stock.AskSize)
	fmt.Printf("FullExchangeName: %s\n", stock.FullExchangeName)
	fmt.Printf("FinancialCurrency: %s\n", stock.FinancialCurrency)
	fmt.Printf("RegularMarketOpen: %.2f\n", stock.RegularMarketOpen)
	fmt.Printf("AverageDailyVolume3Month: %d\n", stock.AverageDailyVolume3Month)
	fmt.Printf("AverageDailyVolume10Day: %d\n", stock.AverageDailyVolume10Day)
	fmt.Printf("FiftyTwoWeekLowChange: %.2f\n", stock.FiftyTwoWeekLowChange)
	fmt.Printf("FiftyTwoWeekLowChangePercent: %.2f\n", stock.FiftyTwoWeekLowChangePercent)
	fmt.Printf("FiftyTwoWeekRange: %s\n", stock.FiftyTwoWeekRange)
	fmt.Printf("FiftyTwoWeekHighChange: %.2f\n", stock.FiftyTwoWeekHighChange)
	fmt.Printf("FiftyTwoWeekHighChangePercent: %.2f\n", stock.FiftyTwoWeekHighChangePercent)
	fmt.Printf("FiftyTwoWeekLow: %.2f\n", stock.FiftyTwoWeekLow)
	fmt.Printf("FiftyTwoWeekHigh: %.2f\n", stock.FiftyTwoWeekHigh)
	fmt.Printf("FiftyTwoWeekChangePercent: %.2f\n", stock.FiftyTwoWeekChangePercent)
	fmt.Printf("EarningsTimestamp: %d\n", stock.EarningsTimestamp)
	fmt.Printf("EarningsTimestampStart: %d\n", stock.EarningsTimestampStart)
	fmt.Printf("EarningsTimestampEnd: %d\n", stock.EarningsTimestampEnd)
	fmt.Printf("TrailingAnnualDividendRate: %.2f\n", stock.TrailingAnnualDividendRate)
	fmt.Printf("TrailingPE: %.2f\n", stock.TrailingPE)
	fmt.Printf("TrailingAnnualDividendYield: %.2f\n", stock.TrailingAnnualDividendYield)
	fmt.Printf("EpsTrailingTwelveMonths: %.2f\n", stock.EpsTrailingTwelveMonths)
	fmt.Printf("EpsForward: %.2f\n", stock.EpsForward)
	fmt.Printf("EpsCurrentYear: %.2f\n", stock.EpsCurrentYear)
	fmt.Printf("PriceEpsCurrentYear: %.2f\n", stock.PriceEpsCurrentYear)
	fmt.Printf("SharesOutstanding: %d\n", stock.SharesOutstanding)
	fmt.Printf("BookValue: %.2f\n", stock.BookValue)
	fmt.Printf("FiftyDayAverage: %.2f\n", stock.FiftyDayAverage)
	fmt.Printf("FiftyDayAverageChange: %.2f\n", stock.FiftyDayAverageChange)
	fmt.Printf("FiftyDayAverageChangePercent: %.2f\n", stock.FiftyDayAverageChangePercent)
	fmt.Printf("TwoHundredDayAverage: %.2f\n", stock.TwoHundredDayAverage)
	fmt.Printf("TwoHundredDayAverageChange: %.2f\n", stock.TwoHundredDayAverageChange)
	fmt.Printf("TwoHundredDayAverageChangePercent: %.2f\n", stock.TwoHundredDayAverageChangePercent)
	fmt.Printf("MarketCap: %d\n", stock.MarketCap)
	fmt.Printf("ForwardPE: %.2f\n", stock.ForwardPE)
	fmt.Printf("PriceToBook: %.2f\n", stock.PriceToBook)
	fmt.Printf("SourceInterval: %d\n", stock.SourceInterval)
	fmt.Printf("ExchangeDataDelayedBy: %d\n", stock.ExchangeDataDelayedBy)
	fmt.Printf("AverageAnalystRating: %s\n", stock.AverageAnalystRating)
	fmt.Printf("Tradeable: %v\n", stock.Tradeable)
	fmt.Printf("CryptoTradeable: %v\n", stock.CryptoTradeable)
	fmt.Printf("DisplayName: %s\n", stock.DisplayName)
	fmt.Printf("Ticker: %s\n", stock.Ticker)
}

// If you want a constant stream of data on the quote you can use the Stream method on the quote
// An example of this method is contained here
// https://github.com/torbenconto/plutus/blob/master/examples/Stock_Data_Stream/main.go
