package tests

import (
	"net/http"
	"net/http/httptest"
	"reflect"
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

	stock, err := quote.NewQuote("GOOG", server.URL)
	if err != nil {
		t.Error("Error fetching data for quote", err)
	}

	for _, tc := range quoteTestCases {
		if fieldValue := getField(stock, tc.field); fieldValue != tc.value {
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

	stock, err := quote.NewQuote("GOOG", server.URL)
	if err != nil {
		t.Error("Error fetching data for quote", err)
	}

	delay := time.Second * 1
	stream := stock.Stream(delay)

	receivedStock := <-stream

	for _, tc := range quoteTestCases {
		if fieldValue := getField(receivedStock, tc.field); fieldValue != tc.value {
			t.Errorf("Expected %s to be %v, got %v", tc.field, tc.value, fieldValue)
		}
	}
}

func getField(s interface{}, field string) interface{} {
	r := reflect.ValueOf(s)
	f := reflect.Indirect(r).FieldByName(field)
	return f.Interface()
}
