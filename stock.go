package plutus

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/gocolly/colly/v2"
)

type Stock struct {
	Ticker                  string
	Price                   float64
	ChangePrice             float64
	ChangePercent           float64
	PrevClose               float64
	OpenPrice               float64
	BidPrice                float64
	AskPrice                float64
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
}

// NewStock creates a new Stock instance for the given ticker.
func NewStock(ticker string) (*Stock, error) {
	c := colly.NewCollector()

	stock := &Stock{
		Ticker:    ticker,
		Collector: c,
	}

	return stock.Populate()
}

// Populate fills in the fields of the Stock struct with data scraped from Yahoo Finance.
func (s *Stock) Populate() (*Stock, error) {
	var err error

	url := fmt.Sprintf("https://finance.yahoo.com/quote/%s", s.Ticker)

	s.Collector.OnHTML("fin-streamer", func(h *colly.HTMLElement) {
		switch h.Attr("data-field") {
		case "regularMarketPrice", "preMarketPrice", "postMarketPrice":
			if isPrimary(h.Attr("active")) {
				s.Price, _ = strconv.ParseFloat(h.Text, 64)
			}
		case "regularMarketChange", "preMarketChange", "postMarketChange":
			if isPrimary(h.Attr("active")) {
				chng, _ := strconv.ParseFloat(h.Text, 64)
				s.ChangePrice = chng
			}
		case "regularMarketChangePercent", "preMarketChangePercent", "postMarketChangePercent":
			if isPrimary(h.Attr("active")) {
				percentString := cleanPercentage(h.Text)
				percentFloat, _ := strconv.ParseFloat(percentString, 64)
				s.ChangePercent = percentFloat
			}
		}
	})

	s.Collector.OnHTML("tr", func(h *colly.HTMLElement) {
		var values []string

		h.ForEach("td", func(i int, t *colly.HTMLElement) {
			values = append(values, t.Text)

			if len(values) == 2 {
				s.setField(YFTableMap[values[0]], values[1])
				values = nil
			}
		})
	})

	s.Collector.OnError(func(r *colly.Response, e error) {
		err = fmt.Errorf("HTTP request error: %v", e)
	})

	err = s.Collector.Visit(url)
	if err != nil {
		return nil, fmt.Errorf("error scraping data: %v", err)
	}

	return s, nil
}

// Helper function to set the struct field based on its type.
func (s *Stock) setField(fieldName string, value string) {
	val := reflect.ValueOf(s).Elem()
	field := val.FieldByName(fieldName)

	switch field.Kind() {
	case reflect.String:
		field.SetString(value)
	case reflect.Float64:
		fieldFloat, _ := strconv.ParseFloat(value, 64)
		field.SetFloat(fieldFloat)
	case reflect.Int:
		fieldInt, _ := strconv.Atoi(value)
		field.SetInt(int64(fieldInt))
	}
}
