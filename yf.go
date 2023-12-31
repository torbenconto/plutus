package plutus

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gocolly/colly/v2"
)

// Internal struct for yf provider
type p_YahooFinanceProvider struct{}

// Expose variable containing provider
var YahooFinanceProvider *p_YahooFinanceProvider = &p_YahooFinanceProvider{}

// Populate fills in the fields of the Stock struct with data scraped from Yahoo Finance.
func (p *p_YahooFinanceProvider) Populate(s *Stock, apiKey ...string) (*Stock, error) {
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
				percentString := cleanNumber(h.Text)
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
				if values[0] == "Day's Range" || values[0] == "52 Week Range" {
					parts := strings.Split(values[1], "-")

					lowHalf := strings.TrimSpace(parts[0])
					highHalf := strings.TrimSpace(parts[1])

					low, _ := strconv.ParseFloat(lowHalf, 64)
					high, _ := strconv.ParseFloat(highHalf, 64)

					if values[0] == "52 Week Range" {
						s.FiftyTwoWeekLow = low
						s.FiftyTwoWeekHigh = high
					} else {
						s.DayLow = low
						s.DayHigh = high
					}
				}

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
