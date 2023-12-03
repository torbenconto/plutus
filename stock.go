package main

import (
	"fmt"
	"reflect"

	"github.com/gocolly/colly/v2"
)

type Stock struct {
	Ticker                  string
	Price                   float64
	Change                  float64
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
}

func NewStock(Ticker string) (*Stock, error) {
	stock := &Stock{
		Ticker: Ticker,
	}

	return stock.Populate()
}

func (s *Stock) Populate() (*Stock, error) {
	// New colly scraper
	c := colly.NewCollector()

	var err error

	// Format url string
	url := fmt.Sprintf("https://finance.yahoo.com/quote/%s", s.Ticker)

	c.OnHTML("fin-streamer", func(h *colly.HTMLElement) {
		switch h.Attr("data-field") {
		case "regularMarketPrice":
			if isPrimary(h.Attr("active")) {
				s.Price = parseNumber(h.Text)
			}

		case "regularMarketChange":
			if isPrimary(h.Attr("active")) {
				s.Change = parseNumber(h.Text)
			}
		}
	})

	c.OnHTML("tr", func(h *colly.HTMLElement) {
		var values []string

		h.ForEach("td", func(i int, t *colly.HTMLElement) {
			text := t.Text
			values = append(values, text)

			if len(values) == 2 {
				val := reflect.ValueOf(s).Elem()

				field := val.FieldByName(YFTableMap[values[0]])

				if field.Kind() == reflect.String {
					field.SetString(values[1])
				} else if field.Kind() == reflect.Float64 {
					num := parseNumber(values[1])
					field.SetFloat(num)
				} else if field.Kind() == reflect.Int {
					num := parseInt(values[1])
					field.SetInt(int64(num))
				}

				values = nil
			}
		})
	})

	c.OnError(func(r *colly.Response, e error) {
		err = fmt.Errorf("error making HTTP request: %v", e)
	})

	err = c.Visit(url)
	if err != nil {
		return nil, fmt.Errorf("error scraping data: %v", err)
	}

	return s, nil
}

func parseNumber(s string) float64 {
	var result float64
	fmt.Sscanf(s, "%f", &result)
	return result
}

func parseInt(s string) int {
	var result int
	fmt.Sscanf(s, "%d", &result)
	return result
}
