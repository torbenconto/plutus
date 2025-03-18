package plutus

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/torbenconto/plutus/interval"
	_range "github.com/torbenconto/plutus/range"
)

type HistoricalPricePoint struct {
	Time   int64
	Open   float64
	Close  float64
	High   float64
	Low    float64
	Volume int64
}

type historicalQuoteResponse struct {
	Chart struct {
		Result []struct {
			Indicators struct {
				Quote []struct {
					Close  []float64 `json:"close"`
					Open   []float64 `json:"open"`
					Volume []int64   `json:"volume"`
					High   []float64 `json:"high"`
					Low    []float64 `json:"low"`
				} `json:"quote"`
			} `json:"indicators"`
			Timestamp []int64 `json:"timestamp"`
		} `json:"result"`
		Error map[string]string `json:"error"`
	} `json:"chart"`
}

/*
GetHistoricalQuote returns a slice of HistoricalPricePoint structs for the given symbol, range, and interval. If the symbol is not found, an error is returned.
The url parameter is optional and is used for testing purposes.
*/
func GetHistoricalQuote(symbol string, r _range.Range, i interval.Interval, url ...string) (*[]HistoricalPricePoint, error) {
	var usedUrl string
	if len(url) > 0 {
		usedUrl = url[0]
	} else {
		usedUrl = historicalUrl(symbol, r.String(), i.String())
	}

	req, err := http.NewRequest(http.MethodGet, usedUrl, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", useragent)
	req.Header.Set("Cookie", cookie)

	get, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer get.Body.Close()

	if get.StatusCode != http.StatusOK {
		return nil, err
	}

	body, err := io.ReadAll(get.Body)
	if err != nil {
		return nil, err
	}

	var hqr historicalQuoteResponse
	if err := json.Unmarshal(body, &hqr); err != nil {
		return nil, err
	}

	if len(hqr.Chart.Error) > 0 {
		return nil, fmt.Errorf("error returned from API: %s, %s", hqr.Chart.Error["code"], hqr.Chart.Error["description"])
	}

	if len(hqr.Chart.Result) == 0 {
		return nil, fmt.Errorf("no results returned from API")
	}

	quote := hqr.Chart.Result[0].Indicators.Quote[0]
	times, closes, opens, volumes, highs, lows := hqr.Chart.Result[0].Timestamp, quote.Close, quote.Open, quote.Volume, quote.High, quote.Low

	var historicalPricePoints []HistoricalPricePoint
	for i := 0; i < len(times); i++ {
		historicalPricePoints = append(historicalPricePoints, HistoricalPricePoint{
			Time:   times[i],
			Open:   opens[i],
			Close:  closes[i],
			High:   highs[i],
			Low:    lows[i],
			Volume: volumes[i],
		})
	}

	return &historicalPricePoints, nil
}
