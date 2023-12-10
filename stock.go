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

// Fill in the fields of the Stock struct with data scraped from yahoo finance
func (s *Stock) Populate() (*Stock, error) {
	// New colly scraper
	c := colly.NewCollector()

	var err error

	// Format url string
	url := fmt.Sprintf("https://finance.yahoo.com/quote/%s", s.Ticker)

	// Loop over <fin-streamer> elements, on yahoo finance these contain price data for the stock
	c.OnHTML("fin-streamer", func(h *colly.HTMLElement) {
		switch h.Attr("data-field") {
		// If data-field = "regularMarketPrice"
		case "regularMarketPrice":
			// Check if the price belongs to the main stock we want data on, not Dow Jones or SPY or something
			if isPrimary(h.Attr("active")) {
				// Parse float into S.price
				s.Price, _ = strconv.ParseFloat(h.Text, 64)
			}

		// If data-field = "regularMarketPrice"
		case "regularMarketChange":
			if isPrimary(h.Attr("active")) {
				// Parse float of price from text
				chng, _ := strconv.ParseFloat(h.Text, 64)

				// Set ChangePrice field
				s.ChangePrice = chng
			}

		case "regularMarketChangePercent":
			if isPrimary(h.Attr("active")) {

				replacer := strings.NewReplacer("%", "", "+", "", "-", "", "(", "", ")", "")
				percentString := replacer.Replace(h.Text)

				fmt.Println(percentString)

				percentFloat, _ := strconv.ParseFloat(percentString, 64)

				s.ChangePercent = percentFloat
			}
		}
	})

	// Loop over rows of the table, on yahoo finance this contains extra data about the stock
	c.OnHTML("tr", func(h *colly.HTMLElement) {
		// Create values array
		var values []string

		// For each table item
		h.ForEach("td", func(i int, t *colly.HTMLElement) {
			text := t.Text
			values = append(values, text)

			// On yf there are 2 tds of importance in each tr, the first one has the title of the data and the second one has the data, by checking if the length of the values is 2, we bundle them together
			if len(values) == 2 {
				val := reflect.ValueOf(s).Elem()

				// Translate the imperfect name of the title field into the name of the struct
				field := val.FieldByName(YFTableMap[values[0]])

				// Switch the type of the data to make sure types match for insertion into the struct
				switch field.Kind() {
				case reflect.String:
					field.SetString(values[1])
				case reflect.Float64:
					fieldFloat, _ := strconv.ParseFloat(values[1], 64)
					field.SetFloat(fieldFloat)
				case reflect.Int:
					fieldInt, _ := strconv.Atoi(values[1])
					field.SetInt(int64(fieldInt))
				}

				// reset values array to bundle the next 2 tds
				values = nil
			}
		})
	})

	// Self explanatory
	c.OnError(func(r *colly.Response, e error) {
		err = fmt.Errorf("error making HTTP request: %v", e)
	})

	err = c.Visit(url)
	if err != nil {
		return nil, fmt.Errorf("error scraping data: %v", err)
	}

	return s, nil
}
