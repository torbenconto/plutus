package quote

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"github.com/torbenconto/plutus/internal/table"
	"github.com/torbenconto/plutus/internal/util"
	"reflect"
	"strconv"
	"strings"
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
			if util.IsPrimary(h.Attr("active")) {
				s.Price, _ = strconv.ParseFloat(h.Text, 64)
			}
		case "regularMarketChange", "preMarketChange", "postMarketChange":
			if util.IsPrimary(h.Attr("active")) {
				chng, _ := strconv.ParseFloat(h.Text, 64)
				s.ChangePrice = chng
			}
		case "regularMarketChangePercent", "preMarketChangePercent", "postMarketChangePercent":
			if util.IsPrimary(h.Attr("active")) {
				percentString := util.CleanNumber(h.Text)
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

				s.setField(table.YFTableMap[values[0]], values[1])
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

	value = util.CleanNumber(value)

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
