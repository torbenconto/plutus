package historical

import (
	"encoding/json"
	"fmt"
	"github.com/torbenconto/plutus"
	"github.com/torbenconto/plutus/config"
	"github.com/torbenconto/plutus/internal/util"
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
	Data   []PricePoint
	Config config.Config
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

func NewHistorical(ticker string, dateRange _range.Range, interval interval.Interval, historicalConfig ...config.Config) (*Historical, error) {
	historical := &Historical{
		Ticker:   ticker,
		Range:    dateRange,
		Interval: interval,
	}

	if len(historicalConfig) > 0 {
		historical.Config = historicalConfig[0]
	} else {
		historical.Config = config.Config{
			Url:       url,
			UserAgent: plutus.UserAgent,
			Cookie:    plutus.Cookie,
		}
	}

	return historical.Populate()
}

func (h *Historical) Populate() (*Historical, error) {
	var req *http.Request
	var err error

	req, err = util.BuildRequestFromConfig(req, h.Config, url, fmt.Sprintf(url, h.Ticker, h.Range, h.Interval))
	if err != nil {
		return nil, fmt.Errorf("error building request: %v", err)
	}

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
