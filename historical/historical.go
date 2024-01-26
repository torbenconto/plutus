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
	Data []TimePricePair
}

type Spark struct {
	Spark struct {
		Result []struct {
			Response []struct {
				Indicators struct {
					Quote []struct {
						Close []float64 `json:"close"`
					} `json:"quote"`
				} `json:"indicators"`
				Timestamp []int64 `json:"timestamp"`
			} `json:"response"`
		} `json:"result"`
	} `json:"spark"`
}

type TimePricePair struct {
	Time  int64
	Price float64
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
	req, err := http.NewRequest("GET", fmt.Sprintf("https://query1.finance.yahoo.com/v7/finance/spark?symbols=%s&range=%s&interval=%s&indicators=close&includeTimestamps=false&includePrePost=false&corsDomain=finance.yahoo.com&.tsrc=finance", h.Ticker, h.Range, h.Interval), nil)

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

	var sparkResponse Spark
	if err := json.Unmarshal(body, &sparkResponse); err != nil {
		fmt.Println("Error:", err)
	}

	priceList := sparkResponse.Spark.Result[0].Response[0].Indicators.Quote[0].Close
	timeList := sparkResponse.Spark.Result[0].Response[0].Timestamp

	tuples := make([]TimePricePair, len(priceList))

	for _, i := range priceList {
		for _, t := range timeList {
			tuples = append(tuples, TimePricePair{Time: t, Price: i})
		}
	}

	h.Data = tuples

	return h, nil
}
