package tests

import (
	"github.com/torbenconto/plutus/config"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/torbenconto/plutus/quote"
)

func TestQuote(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write(quoteServerData)
		if err != nil {
			t.Error("Error writing response")
		}
	}))
	defer server.Close()

	stock, err := quote.NewQuote("GOOG", config.Config{
		Url: server.URL,
	})
	if err != nil {
		t.Error("Error fetching data for quote", err)
	}

	for _, tc := range quoteTestCases {
		if fieldValue := GetField(stock, tc.field); fieldValue != tc.value {
			t.Errorf("Expected %s to be %v, got %v", tc.field, tc.value, fieldValue)
		}
	}
}

func TestYahooQuoteApi(t *testing.T) {
	stock, err := quote.NewQuote("GOOG")
	if err != nil {
		t.Error("Error fetching data for quote", err)
	}

	if stock == nil {
		t.Error("Stock is nil")
	}
}

func TestQuoteStream(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write(quoteServerData)
		if err != nil {
			t.Error("Error writing response")
		}
	}))
	defer server.Close()

	stock, err := quote.NewQuote("GOOG", config.Config{
		Url: server.URL,
	})
	if err != nil {
		t.Error("Error fetching data for quote", err)
	}

	delay := time.Second * 1
	stream := stock.Stream(delay)

	receivedStock := <-stream

	for _, tc := range quoteTestCases {
		if fieldValue := GetField(receivedStock, tc.field); fieldValue != tc.value {
			t.Errorf("Expected %s to be %v, got %v", tc.field, tc.value, fieldValue)
		}
	}
}
