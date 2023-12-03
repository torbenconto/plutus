package main

import (
	"fmt"
	"strconv"

	"github.com/gocolly/colly/v2"
)

type Stock struct {
	Ticker string
	Price  float64
	Change float64
}

func NewStock(Ticker string) (*Stock, error) {
	stock := &Stock{
		Ticker: Ticker,
	}

	return stock.Populate()
}

func (s *Stock) Populate() (*Stock, error) {
	c := colly.NewCollector()

	var err error

	url := fmt.Sprintf("https://finance.yahoo.com/quote/%s", s.Ticker)

	c.OnHTML("fin-streamer", func(h *colly.HTMLElement) {
		if h.Attr("data-field") == "regularMarketPrice" && h.Attr("active") == "" {
			s.Price, err = strconv.ParseFloat(h.Text, 64)
			if err != nil {
				err = fmt.Errorf("error parsing price: %v", err)
			}
		}
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
