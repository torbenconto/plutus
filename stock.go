package plutus

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gocolly/colly/v2"
)

type Stock struct {
	Ticker                     string
	Price                      float64
	ChangePrice                float64
	ChangePercent              float64
	PrevClose                  float64
	OpenPrice                  float64
	BidPrice                   string
	AskPrice                   string
	DayHigh                    float64
	DayLow                     float64
	FiftyTwoWeekLow            float64
	FiftyTwoWeekHigh           float64
	Volume                     int
	AvgVolume                  int
	MarketCap                  string
	Beta                       float64
	PE                         float64
	EPS                        float64
	FiftyDayMovingAverage      float64
	TwoHundredDayMovingAverage float64
	SharesOutstanding          float64
	EarningsDate               string
	ForwardDividendAndYield    string
	DividendDate               string
	ExDividendDate             string
	OneYearTargetEst           float64
	Collector                  *colly.Collector
}

// NewStock creates a new Stock instance for the given ticker.
func NewStock(ticker string) (*Stock, error) {
	c := colly.NewCollector()

	stock := &Stock{
		Ticker:    strings.ToUpper(ticker),
		Collector: c,
	}

	return stock.Populate()
}

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
