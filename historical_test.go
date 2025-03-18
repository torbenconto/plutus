package plutus_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/torbenconto/plutus/v2"
	"github.com/torbenconto/plutus/v2/interval"
	_range "github.com/torbenconto/plutus/v2/range"
)

var historicalServerData = []byte(`
	{
  "chart": {
    "result": [
      {
        "indicators": {
          "quote": [
            {
              "close": [
                12.0,
                14.6,
                13.2
              ],
              "open": [
                16.0,
                7.4,
                9.62
              ],
              "volume": [
                212010,
                43021,
                350511
              ],
              "high": [
                13.6,
                15.2,
                7.2
              ],
              "low": [
                10.0,
                12.2,
                2.31
              ]
            }
          ]
        },
        "timestamp": [
          1615891200,
          1615977600,
          1616064000
        ]
      }
    ]
  }
}`)

var historicalTestCases = []plutus.HistoricalPricePoint{
	{1615891200, 16.0, 12.0, 13.6, 10.0, 212010},
	{1615977600, 7.4, 14.6, 15.2, 12.2, 43021},
	{1616064000, 9.62, 13.2, 7.2, 2.31, 350511},
}

func TestHistoricalPopulate(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write(historicalServerData)
		if err != nil {
			t.Error("Error writing response")
		}
	}))
	defer server.Close()

	data, err := plutus.GetHistoricalQuote("GOOG", _range.OneDay, interval.OneMinute, server.URL)
	if err != nil {
		t.Error("Error fetching data for historical", err)
	}

	if len(*data) != 3 {
		t.Errorf("Expected 3 data points, got %d", len(*data))
	}

	for i, tc := range historicalTestCases {
		if (*data)[i].Time != tc.Time {
			t.Errorf("Expected time to be %d, got %d", tc.Time, (*data)[i].Time)
		}
		if (*data)[i].Open != tc.Open {
			t.Errorf("Expected open to be %f, got %f", tc.Open, (*data)[i].Open)
		}
		if (*data)[i].Close != tc.Close {
			t.Errorf("Expected close to be %f, got %f", tc.Close, (*data)[i].Close)
		}
		if (*data)[i].High != tc.High {
			t.Errorf("Expected high to be %f, got %f", tc.High, (*data)[i].High)
		}
		if (*data)[i].Low != tc.Low {
			t.Errorf("Expected low to be %f, got %f", tc.Low, (*data)[i].Low)
		}
		if (*data)[i].Volume != tc.Volume {
			t.Errorf("Expected volume to be %d, got %d", tc.Volume, (*data)[i].Volume)
		}
	}

}

func TestYahooHistoricalApi(t *testing.T) {
	data, err := plutus.GetHistoricalQuote("GOOG", _range.OneDay, interval.OneMinute)
	if err != nil {
		t.Error("Error fetching data for historical", err)
	}

	if len(*data) == 0 {
		t.Error("Data is empty")
	}
}
