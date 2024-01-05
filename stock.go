package plutus

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/gocolly/colly/v2"
)

type Stock struct {
	Ticker                  string
	Price                   float64
	ChangePrice             float64
	ChangePercent           float64
	PrevClose               float64
	OpenPrice               float64
	BidPrice                string
	AskPrice                string
	DayRange                string
	FiftyTwoWeekRange       string
	Volume                  int
	AvgVolume               int
	MarketCap               string
	Beta                    float64
	PE                      float64
	EPS                     float64
	EarningsDate            string
	ForwardDividendAndYield string
	ExDividendDate          string
	OneYearTargetEst        float64
	Collector               *colly.Collector
	Provider                StockDataProvider
}

// NewStock creates a new Stock instance for the given ticker.
func NewStock(ticker string, provider StockDataProvider, apiKey ...string) (*Stock, error) {
	c := colly.NewCollector()

	stock := &Stock{
		Ticker:    strings.ToUpper(ticker),
		Collector: c,
		Provider:  provider,
	}

	return stock.Provider.Populate(stock, apiKey...)
}

// Helper function to set the struct field based on its type.
func (s *Stock) setField(fieldName string, value string) {
	val := reflect.ValueOf(s).Elem()
	field := val.FieldByName(fieldName)

	value = cleanNumber(value)

	switch field.Kind() {
	case reflect.String:
		field.SetString(value)
	case reflect.Float64:
		fmt.Println(fieldName, value)
		fieldFloat, _ := strconv.ParseFloat(value, 64)
		field.SetFloat(fieldFloat)
	case reflect.Int:
		fieldInt, _ := strconv.Atoi(value)
		field.SetInt(int64(fieldInt))
	}
}
