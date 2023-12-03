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
	// New colly scraper
	c := colly.NewCollector()

	var err error

	// Format url string
	url := fmt.Sprintf("https://finance.yahoo.com/quote/%s", s.Ticker)

	c.OnHTML("fin-streamer", func(h *colly.HTMLElement) {
		switch h.Attr("data-field") {
		case "regularMarketPrice":
			if isPrimary(h.Attr("active")) {
				s.Price, err = strconv.ParseFloat(h.Text, 64)
				if err != nil {
					err = fmt.Errorf("error parsing price: %v", err)
				}
			}

		case "regularMarketChange":
			if isPrimary(h.Attr("active")) {
				s.Change, err = strconv.ParseFloat(h.Text, 64)
				if err != nil {
					err = fmt.Errorf("error parsing change: %v", err)
				}
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
