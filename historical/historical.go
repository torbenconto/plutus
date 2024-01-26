package historical

import (
	"encoding/json"
	"fmt"
	"github.com/torbenconto/plutus"
	"github.com/torbenconto/plutus/interval"
	"github.com/torbenconto/plutus/range"
	"io"
	"net/http"
)

type Historical struct {
	Ticker   string
	Range    _range.Range
	Interval interval.Interval
	// Set of structs containing time and price data paired together
	Data []PricePoint
}

type Response struct {
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
	} `json:"chart"`
}

type PricePoint struct {
	Time   int64
	Open   float64
	Close  float64
	High   float64
	Low    float64
	Volume int64
}

func NewHistorical(ticker string, dateRange _range.Range, interval interval.Interval) (*Historical, error) {
	historical := &Historical{
		Ticker:   ticker,
		Range:    dateRange,
		Interval: interval,
	}

	return historical.Populate()
}

func (h *Historical) Populate() (*Historical, error) {
	var err error

	// Get quote
	req, err := http.NewRequest("GET", fmt.Sprintf("https://query1.finance.yahoo.com/v8/finance/chart/%s?region=US&lang=en-US&includePrePost=false&range=%s&interval=%s&useYfid=true&corsDomain=finance.yahoo.com&.tsrc=finance&indicators=quote", h.Ticker, h.Range, h.Interval), nil)

	req.Header.Set("User-Agent", plutus.UserAgent)
	req.Header.Set("Cookie", plutus.Cookie)

	get, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %v", err)
	}

	defer get.Body.Close()

	body, err := io.ReadAll(get.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %v", err)
	}

	var chartResponse Response
	if err := json.Unmarshal(body, &chartResponse); err != nil {
		fmt.Println("Error:", err)
	}

	closeList := chartResponse.Chart.Result[0].Indicators.Quote[0].Close
	openList := chartResponse.Chart.Result[0].Indicators.Quote[0].Open
	volumeList := chartResponse.Chart.Result[0].Indicators.Quote[0].Volume
	highList := chartResponse.Chart.Result[0].Indicators.Quote[0].High
	lowList := chartResponse.Chart.Result[0].Indicators.Quote[0].Low
	timeList := chartResponse.Chart.Result[0].Timestamp

	tuples := make([]PricePoint, 0)

	for index, time := range timeList {
		tuples = append(tuples, PricePoint{Time: time, Open: openList[index], Close: closeList[index], High: highList[index], Low: lowList[index], Volume: volumeList[index]})
	}

	h.Data = tuples

	return h, nil
}
