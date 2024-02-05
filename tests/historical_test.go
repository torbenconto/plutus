package main

import (
	"github.com/torbenconto/plutus/historical"
	"github.com/torbenconto/plutus/interval"
	_range "github.com/torbenconto/plutus/range"
	"net/http"
	"net/http/httptest"
	"testing"
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

var historicalTestCases = []struct {
	Time   int64
	Open   float64
	Close  float64
	High   float64
	Low    float64
	Volume int64
}{
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

	historical.URL = server.URL

	data, err := historical.NewHistorical("GOOG", _range.OneDay, interval.OneMinute)
	if err != nil {
		t.Error("Error fetching data for historical", err)
	}

	if len(data.Data) != 3 {
		t.Errorf("Expected 3 data points, got %d", len(data.Data))
	}

	for i, tc := range historicalTestCases {
		if data.Data[i].Time != tc.Time {
			t.Errorf("Expected time to be %d, got %d", tc.Time, data.Data[i].Time)
		}
		if data.Data[i].Open != tc.Open {
			t.Errorf("Expected open to be %f, got %f", tc.Open, data.Data[i].Open)
		}
		if data.Data[i].Close != tc.Close {
			t.Errorf("Expected close to be %f, got %f", tc.Close, data.Data[i].Close)
		}
		if data.Data[i].High != tc.High {
			t.Errorf("Expected high to be %f, got %f", tc.High, data.Data[i].High)
		}
		if data.Data[i].Low != tc.Low {
			t.Errorf("Expected low to be %f, got %f", tc.Low, data.Data[i].Low)
		}
		if data.Data[i].Volume != tc.Volume {
			t.Errorf("Expected volume to be %d, got %d", tc.Volume, data.Data[i].Volume)
		}
	}

}
