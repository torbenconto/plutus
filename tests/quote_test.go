package main

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	"github.com/torbenconto/plutus/quote"
)

var quoteServerData = []byte(`{
   "quoteResponse": {
     "result": [{
       "symbol": "GOOG",
       "regularMarketPrice": 1234.56,
       "regularMarketChangePercent": 0.1234,
       "bidSize": 100,
       "askSize": 200,
       "fullExchangeName": "NASDAQ",
       "financialCurrency": "USD",
       "regularMarketOpen": 1200.00,
       "averageDailyVolume3Month": 1500000,
       "averageDailyVolume10Day": 1300000,
       "fiftyTwoWeekLowChange": 234.56,
       "fiftyTwoWeekLowChangePercent": 0.2345,
       "fiftyTwoWeekRange": "1000.00 - 1300.00",
       "fiftyTwoWeekHighChange": 34.56,
       "fiftyTwoWeekHighChangePercent": 0.0345,
       "fiftyTwoWeekLow": 1000.00,
       "fiftyTwoWeekHigh": 1300.00,
       "fiftyTwoWeekChangePercent": 0.30,
       "earningsTimestamp": 1615891200,
       "earningsTimestampStart": 1615891200,
       "earningsTimestampEnd": 1615891200,
       "trailingAnnualDividendRate": 20.00,
       "trailingPE": 30.00,
       "trailingAnnualDividendYield": 0.015,
       "epsTrailingTwelveMonths": 40.00,
       "epsForward": 45.00,
       "epsCurrentYear": 42.00,
       "priceEpsCurrentYear": 30.00,
       "sharesOutstanding": 300000000,
       "bookValue": 50.00,
       "fiftyDayAverage": 1250.00,
       "fiftyDayAverageChange": -15.44,
       "fiftyDayAverageChangePercent": -0.0123,
       "twoHundredDayAverage": 1200.00,
       "twoHundredDayAverageChange": 34.56,
       "twoHundredDayAverageChangePercent": 0.0288,
       "marketCap": 1000000000,
       "forwardPE": 25.00,
       "priceToBook": 24.68,
       "sourceInterval": 15,
       "exchangeDataDelayedBy": 0,
       "averageAnalystRating": "Buy",
       "tradeable": true,
       "cryptoTradeable": false,
       "displayName": "Google Inc."
     }]
   }
 }`)

var quoteTestCases = []struct {
	field string
	value interface{}
}{
	{"Ticker", "GOOG"},
	{"RegularMarketPrice", 1234.56},
	{"BidSize", 100},
	{"RegularMarketChangePercent", 0.1234},
	{"FullExchangeName", "NASDAQ"},
	{"FinancialCurrency", "USD"},
	{"RegularMarketOpen", 1200.00},
	{"AverageDailyVolume3Month", int64(1500000)},
	{"AverageDailyVolume10Day", int64(1300000)},
	{"FiftyTwoWeekLowChange", 234.56},
	{"FiftyTwoWeekLowChangePercent", 0.2345},
	{"FiftyTwoWeekRange", "1000.00 - 1300.00"},
	{"FiftyTwoWeekHighChange", 34.56},
	{"FiftyTwoWeekHighChangePercent", 0.0345},
	{"FiftyTwoWeekLow", 1000.00},
	{"FiftyTwoWeekHigh", 1300.00},
	{"FiftyTwoWeekChangePercent", 0.30},
	{"EarningsTimestamp", int64(1615891200)},
	{"EarningsTimestampStart", int64(1615891200)},
	{"EarningsTimestampEnd", int64(1615891200)},
	{"TrailingAnnualDividendRate", 20.00},
	{"TrailingPE", 30.00},
	{"TrailingAnnualDividendYield", 0.015},
	{"EpsTrailingTwelveMonths", 40.00},
	{"EpsForward", 45.00},
	{"EpsCurrentYear", 42.00},
	{"PriceEpsCurrentYear", 30.00},
	{"SharesOutstanding", int64(300000000)},
	{"BookValue", 50.00},
	{"FiftyDayAverage", 1250.00},
	{"FiftyDayAverageChange", -15.44},
	{"FiftyDayAverageChangePercent", -0.0123},
	{"TwoHundredDayAverage", 1200.00},
	{"TwoHundredDayAverageChange", 34.56},
	{"TwoHundredDayAverageChangePercent", 0.0288},
	{"MarketCap", int64(1000000000)},
	{"ForwardPE", 25.00},
	{"PriceToBook", 24.68},
	{"SourceInterval", 15},
	{"ExchangeDataDelayedBy", 0},
	{"AverageAnalystRating", "Buy"},
	{"Tradeable", true},
	{"CryptoTradeable", false},
	{"DisplayName", "Google Inc."},
}

func TestQuote(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write(quoteServerData)
		if err != nil {
			t.Error("Error writing response")
		}
	}))
	defer server.Close()

	stock, err := quote.NewQuote("GOOG", server.URL)
	if err != nil {
		t.Error("Error fetching data for quote", err)
	}

	for _, tc := range quoteTestCases {
		if fieldValue := getField(stock, tc.field); fieldValue != tc.value {
			t.Errorf("Expected %s to be %v, got %v", tc.field, tc.value, fieldValue)
		}
	}
}

func TestYahooQuoteApi(t *testing.T) {
	stock, err := quote.NewQuote("GOOG")
	if err != nil {
		t.Error("Error fetching data for quote", err)
	}

	if stock == nil {
		t.Error("Stock is nil")
	}
}

func TestQuoteStream(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write(quoteServerData)
		if err != nil {
			t.Error("Error writing response")
		}
	}))
	defer server.Close()

	stock, err := quote.NewQuote("GOOG", server.URL)
	if err != nil {
		t.Error("Error fetching data for quote", err)
	}

	delay := time.Second * 1
	stream := stock.Stream(delay)

	receivedStock := <-stream

	for _, tc := range quoteTestCases {
		if fieldValue := getField(receivedStock, tc.field); fieldValue != tc.value {
			t.Errorf("Expected %s to be %v, got %v", tc.field, tc.value, fieldValue)
		}
	}
}

func getField(s interface{}, field string) interface{} {
	r := reflect.ValueOf(s)
	f := reflect.Indirect(r).FieldByName(field)
	return f.Interface()
}
