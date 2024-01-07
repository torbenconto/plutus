package plutus

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

type p_AlphaVantageApiProvider struct {
	apiKey string
}

var AlphaVantageApiProvider *p_AlphaVantageApiProvider = &p_AlphaVantageApiProvider{}

type Overview struct {
	MarketCap                  string `json:"MarketCapitalization"`
	PERatio                    string `json:"PERatio"`
	EPS                        string `json:"EPS"`
	Beta                       string `json:"Beta"`
	FiftyTwoWeekHigh           string `json:"52WeekHigh"`
	FiftyTwoWeekLow            string `json:"52WeekLow"`
	FiftyDayMovingAverage      string `json:"50DayMovingAverage"`
	TwoHundredDayMovingAverage string `json:"200DayMovingAverage"`
	SharesOutstanding          string `json:"SharesOutstanding"`
	DividendDate               string `json:"DividendDate"`
	ExDividendDate             string `json:"ExDividendDate"`
}

type Quote struct {
	GlobalQuote struct {
		Price         string `json:"05. price"`
		ChangePrice   string `json:"09. change"`
		ChangePercent string `json:"10. change percent"`
		OpenPrice     string `json:"02. open"`
		DayHigh       string `json:"03. high"`
		DayLow        string `json:"04. low"`
		Volume        string `json:"06. volume"`
	} `json:"Global Quote"`
}

func (p *p_AlphaVantageApiProvider) Populate(s *Stock, apiKey ...string) (*Stock, error) {
	if len(apiKey) > 0 {
		p.apiKey = apiKey[0]
	} else {
		return nil, ErrNoAPIKey
	}

	overviewRequest, err := http.Get("https://www.alphavantage.co/query?function=OVERVIEW&symbol=" + s.Ticker + "&apikey=" + p.apiKey)
	if err != nil {
		return nil, err
	}

	overviewBody, err := io.ReadAll(overviewRequest.Body)
	if err != nil {
		return nil, err
	}

	var overview Overview

	err = json.Unmarshal(overviewBody, &overview)
	if err != nil {
		return s, err
	}

	s.MarketCap = overview.MarketCap
	s.PE = parseFloat(overview.PERatio)
	s.EPS = parseFloat(overview.EPS)
	s.Beta = parseFloat(overview.Beta)
	s.FiftyTwoWeekLow = parseFloat(overview.FiftyTwoWeekLow)
	s.FiftyTwoWeekHigh = parseFloat(overview.FiftyTwoWeekHigh)
	s.FiftyDayMovingAverage = parseFloat(overview.FiftyDayMovingAverage)
	s.TwoHundredDayMovingAverage = parseFloat(overview.TwoHundredDayMovingAverage)
	s.SharesOutstanding = parseFloat(overview.SharesOutstanding)
	s.DividendDate = overview.DividendDate
	s.ExDividendDate = overview.ExDividendDate

	quoteRequest, err := http.Get("https://www.alphavantage.co/query?function=GLOBAL_QUOTE&symbol=" + s.Ticker + "&apikey=" + p.apiKey)
	if err != nil {
		return nil, err
	}

	quoteBody, err := io.ReadAll(quoteRequest.Body)
	if err != nil {
		return nil, err
	}

	var quote Quote

	err = json.Unmarshal(quoteBody, &quote)
	if err != nil {
		return s, err
	}

	s.Price = parseFloat(quote.GlobalQuote.Price)
	s.ChangePrice = parseFloat(quote.GlobalQuote.ChangePrice)
	s.ChangePercent = parseFloat(quote.GlobalQuote.ChangePercent)
	s.OpenPrice = parseFloat(quote.GlobalQuote.OpenPrice)
	s.DayHigh = parseFloat(quote.GlobalQuote.DayHigh)
	s.DayLow = parseFloat(quote.GlobalQuote.DayLow)
	s.Volume, _ = strconv.Atoi(quote.GlobalQuote.Volume)

	return s, nil
}

// parseFloat converts numeric values from strings to float64
func parseFloat(value interface{}) float64 {
	switch v := value.(type) {
	case float64:
		return v
	case string:
		num, err := strconv.ParseFloat(cleanNumber(v), 64)
		if err != nil {
			return 0.0
		}
		return num
	default:
		return 0.0
	}
}
