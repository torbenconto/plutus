package historical

import (
	"encoding/json"
	"fmt"
	"github.com/torbenconto/plutus"
	"github.com/torbenconto/plutus/interval"
	"github.com/torbenconto/plutus/range"
	"io"
	"net/http"
	"strings"
)

type Historical struct {
	Ticker   string
	Range    _range.Range
	Interval interval.Interval
	// Set of structs containing time and price data paired together
	Data []PricePoint
	URL  string
}

type response struct {
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

type PricePoint struct {
	Time   int64
	Open   float64
	Close  float64
	High   float64
	Low    float64
	Volume int64
}

func NewHistorical(ticker string, dateRange _range.Range, interval interval.Interval, apiUrl ...string) (*Historical, error) {
	historical := &Historical{
		Ticker:   ticker,
		Range:    dateRange,
		Interval: interval,
	}

	// apiUrl is mainly used for testing purposes, additional functionality will be added in the future
	if len(apiUrl) > 0 {
		historical.URL = apiUrl[0]
	} else {
		historical.URL = url
	}

	return historical.Populate()
}

func (h *Historical) Populate() (*Historical, error) {
	var req *http.Request
	var err error

	if strings.Count(h.URL, "%s") == 3 {
		req, err = http.NewRequest("GET", fmt.Sprintf(h.URL, h.Ticker, h.Range.String(), h.Interval.String()), nil)
	} else {
		req, err = http.NewRequest("GET", h.URL, nil)
	}

	req.Header.Set("User-Agent", plutus.UserAgent)
	req.Header.Set("Cookie", plutus.Cookie)

	get, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %v", err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("Error:", err)
		}
	}(get.Body)

	body, err := io.ReadAll(get.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %v", err)
	}

	var chartResponse response
	if err := json.Unmarshal(body, &chartResponse); err != nil {
		fmt.Println("Error:", err)
	}

	// Check if the Error field is not nil
	if chartResponse.Chart.Error != nil {
		return nil, fmt.Errorf("error returned from API: %s, %s", chartResponse.Chart.Error["code"], chartResponse.Chart.Error["description"])
	}

	// Check if the Result field is empty
	if len(chartResponse.Chart.Result) == 0 {

		return nil, fmt.Errorf("error returned from API: no result returned")
	}

	quote := chartResponse.Chart.Result[0].Indicators.Quote[0]
	timeList := chartResponse.Chart.Result[0].Timestamp
	closeList := quote.Close
	openList := quote.Open
	volumeList := quote.Volume
	highList := quote.High
	lowList := quote.Low

	tuples := make([]PricePoint, 0)

	for index, time := range timeList {
		tuples = append(tuples, PricePoint{Time: time, Open: openList[index], Close: closeList[index], High: highList[index], Low: lowList[index], Volume: volumeList[index]})
	}

	h.Data = tuples

	return h, nil
}
