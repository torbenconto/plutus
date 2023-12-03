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

	c.OnHTML("fin-streamer", func(h *colly.HTMLElement) {

		if h.Attr("data-field") == "regularMarketPrice" && h.Attr("active") == "" {
			s.Price, err = strconv.ParseFloat(h.Text, 64)
			if err != nil {
				err = fmt.Errorf("error parsing price: %e", err)
			}
		}
	})

	if err != nil {
		return nil, err
	}

	err = c.Visit("https://finance.yahoo.com/quote/amd")

	if err != nil {
		err = fmt.Errorf("error scraping data: %e", err)
		return nil, err
	}

	return s, nil
}
