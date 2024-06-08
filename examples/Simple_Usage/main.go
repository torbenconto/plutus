package main

import (
	"fmt"
	"github.com/torbenconto/plutus/stock"
)

func main() {
	// Create a stock object using a ticker, the struct will be auto populated with data when created, no need to call any other methods
	// Returns an error and a plutus.stock.Quote struct, error details what went wrong with fetching data
	//                  no need to capitalize ticker
	//                              |
	//                              v
	data, err := stock.NewQuote("amd")
	if err != nil {
		fmt.Printf("An error occured: %s\n", err)
	}

	// Print all values defined in the Quote struct
	fmt.Printf("Language: %s\n", data.Language)
	fmt.Printf("Region: %s\n", data.Region)
	fmt.Printf("QuoteType: %s\n", data.QuoteType)
	fmt.Printf("TypeDisp: %s\n", data.TypeDisp)
	fmt.Printf("QuoteSourceName: %s\n", data.QuoteSourceName)
	fmt.Printf("Triggerable: %v\n", data.Triggerable)
	fmt.Printf("CustomPriceAlertConfidence: %s\n", data.CustomPriceAlertConfidence)
	fmt.Printf("Currency: %s\n", data.Currency)
	fmt.Printf("MarketState: %s\n", data.MarketState)
	fmt.Printf("RegularMarketChangePercent: %.2f\n", data.RegularMarketChangePercent)
	fmt.Printf("RegularMarketPrice: %.2f\n", data.RegularMarketPrice)
	fmt.Printf("Exchange: %s\n", data.Exchange)
	fmt.Printf("ShortName: %s\n", data.ShortName)
	fmt.Printf("LongName: %s\n", data.LongName)
	fmt.Printf("MessageBoardID: %s\n", data.MessageBoardID)
	fmt.Printf("ExchangeTimezoneName: %s\n", data.ExchangeTimezoneName)
	fmt.Printf("ExchangeTimezoneShortName: %s\n", data.ExchangeTimezoneShortName)
	fmt.Printf("GmtOffSetMilliseconds: %d\n", data.GmtOffSetMilliseconds)
	fmt.Printf("Market: %s\n", data.Market)
	fmt.Printf("EsgPopulated: %v\n", data.EsgPopulated)
	fmt.Printf("FirstTradeDateMilliseconds: %d\n", data.FirstTradeDateMilliseconds)
	fmt.Printf("PriceHint: %d\n", data.PriceHint)
	fmt.Printf("PostMarketChangePercent: %.2f\n", data.PostMarketChangePercent)
	fmt.Printf("PostMarketTime: %d\n", data.PostMarketTime)
	fmt.Printf("PostMarketPrice: %.2f\n", data.PostMarketPrice)
	fmt.Printf("PostMarketChange: %.2f\n", data.PostMarketChange)
	fmt.Printf("RegularMarketChange: %.2f\n", data.RegularMarketChange)
	fmt.Printf("RegularMarketTime: %d\n", data.RegularMarketTime)
	fmt.Printf("RegularMarketDayHigh: %.2f\n", data.RegularMarketDayHigh)
	fmt.Printf("RegularMarketDayRange: %s\n", data.RegularMarketDayRange)
	fmt.Printf("RegularMarketDayLow: %.2f\n", data.RegularMarketDayLow)
	fmt.Printf("RegularMarketVolume: %d\n", data.RegularMarketVolume)
	fmt.Printf("RegularMarketPreviousClose: %.2f\n", data.RegularMarketPreviousClose)
	fmt.Printf("Bid: %.2f\n", data.Bid)
	fmt.Printf("Ask: %.2f\n", data.Ask)
	fmt.Printf("BidSize: %d\n", data.BidSize)
	fmt.Printf("AskSize: %d\n", data.AskSize)
	fmt.Printf("FullExchangeName: %s\n", data.FullExchangeName)
	fmt.Printf("FinancialCurrency: %s\n", data.FinancialCurrency)
	fmt.Printf("RegularMarketOpen: %.2f\n", data.RegularMarketOpen)
	fmt.Printf("AverageDailyVolume3Month: %d\n", data.AverageDailyVolume3Month)
	fmt.Printf("AverageDailyVolume10Day: %d\n", data.AverageDailyVolume10Day)
	fmt.Printf("FiftyTwoWeekLowChange: %.2f\n", data.FiftyTwoWeekLowChange)
	fmt.Printf("FiftyTwoWeekLowChangePercent: %.2f\n", data.FiftyTwoWeekLowChangePercent)
	fmt.Printf("FiftyTwoWeekRange: %s\n", data.FiftyTwoWeekRange)
	fmt.Printf("FiftyTwoWeekHighChange: %.2f\n", data.FiftyTwoWeekHighChange)
	fmt.Printf("FiftyTwoWeekHighChangePercent: %.2f\n", data.FiftyTwoWeekHighChangePercent)
	fmt.Printf("FiftyTwoWeekLow: %.2f\n", data.FiftyTwoWeekLow)
	fmt.Printf("FiftyTwoWeekHigh: %.2f\n", data.FiftyTwoWeekHigh)
	fmt.Printf("FiftyTwoWeekChangePercent: %.2f\n", data.FiftyTwoWeekChangePercent)
	fmt.Printf("EarningsTimestamp: %d\n", data.EarningsTimestamp)
	fmt.Printf("EarningsTimestampStart: %d\n", data.EarningsTimestampStart)
	fmt.Printf("EarningsTimestampEnd: %d\n", data.EarningsTimestampEnd)
	fmt.Printf("TrailingAnnualDividendRate: %.2f\n", data.TrailingAnnualDividendRate)
	fmt.Printf("TrailingPE: %.2f\n", data.TrailingPE)
	fmt.Printf("TrailingAnnualDividendYield: %.2f\n", data.TrailingAnnualDividendYield)
	fmt.Printf("EpsTrailingTwelveMonths: %.2f\n", data.EpsTrailingTwelveMonths)
	fmt.Printf("EpsForward: %.2f\n", data.EpsForward)
	fmt.Printf("EpsCurrentYear: %.2f\n", data.EpsCurrentYear)
	fmt.Printf("PriceEpsCurrentYear: %.2f\n", data.PriceEpsCurrentYear)
	fmt.Printf("SharesOutstanding: %d\n", data.SharesOutstanding)
	fmt.Printf("BookValue: %.2f\n", data.BookValue)
	fmt.Printf("FiftyDayAverage: %.2f\n", data.FiftyDayAverage)
	fmt.Printf("FiftyDayAverageChange: %.2f\n", data.FiftyDayAverageChange)
	fmt.Printf("FiftyDayAverageChangePercent: %.2f\n", data.FiftyDayAverageChangePercent)
	fmt.Printf("TwoHundredDayAverage: %.2f\n", data.TwoHundredDayAverage)
	fmt.Printf("TwoHundredDayAverageChange: %.2f\n", data.TwoHundredDayAverageChange)
	fmt.Printf("TwoHundredDayAverageChangePercent: %.2f\n", data.TwoHundredDayAverageChangePercent)
	fmt.Printf("MarketCap: %d\n", data.MarketCap)
	fmt.Printf("ForwardPE: %.2f\n", data.ForwardPE)
	fmt.Printf("PriceToBook: %.2f\n", data.PriceToBook)
	fmt.Printf("SourceInterval: %d\n", data.SourceInterval)
	fmt.Printf("ExchangeDataDelayedBy: %d\n", data.ExchangeDataDelayedBy)
	fmt.Printf("AverageAnalystRating: %s\n", data.AverageAnalystRating)
	fmt.Printf("Tradeable: %v\n", data.Tradeable)
	fmt.Printf("CryptoTradeable: %v\n", data.CryptoTradeable)
	fmt.Printf("DisplayName: %s\n", data.DisplayName)
	fmt.Printf("Ticker: %s\n", data.Ticker)
}

// If you want a constant stream of data on the stock you can use the Stream method on the stock
// An example of this method is contained here
// https://github.com/torbenconto/plutus/blob/master/examples/Stock_Data_Stream/main.go
