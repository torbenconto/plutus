package news

import (
	"github.com/torbenconto/plutus/config"
	"net/http"
	"net/http/httptest"
	"testing"
)

var newsServerData = []byte(`{
    "news": [
        {
            "uuid": "1234",
            "ticker": "GOOG",
            "title": "Google Inc. is doing great",
            "publisher": "CNBC",
            "link": "https://www.cnbc.com/1234",
            "providerPublishTime": 1615891200,
            "type": "Article",
            "relatedTickers": [
                "AAPL",
                "MSFT"
            ]
        },
        {
            "uuid": "5678",
            "ticker": "MSFT",
            "title": "Microsoft Inc. is doing terrible",
            "publisher": "CNBC",
            "link": "https://www.cnbc.com/5678",
            "providerPublishTime": 1615891200,
            "type": "Article",
            "relatedTickers": [
                "AAPL",
                "MSFT"
            ]
        }
    ]
}`)

var newsTestCases = []struct {
	field string
	value interface{}
}{
	{"Articles", []Article{
		{
			Ticker:              "GOOG",
			Uuid:                "1234",
			Title:               "Google Inc. is doing great",
			Publisher:           "CNBC",
			Link:                "https://www.cnbc.com/1234",
			ProviderPublishTime: 1615891200,
			Type:                "Article",
			RelatedTickers:      []string{"AAPL", "MSFT"},
		},
		{
			Ticker:              "MSFT",
			Uuid:                "5678",
			Title:               "Microsoft Inc. is doing terrible",
			Publisher:           "CNBC",
			Link:                "https://www.cnbc.com/5678",
			ProviderPublishTime: 1615891300,
			Type:                "Article",
			RelatedTickers:      []string{"AAPL", "MSFT"},
		},
	}},
}

func TestNewsPopulate(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write(newsServerData)
		if err != nil {
			t.Error("Error writing response")
		}
	}))
	defer server.Close()

	data, err := NewNews("GOOG", config.Config{
		Url: server.URL,
	})
	if err != nil {
		t.Error("Error fetching data for news", err)
	}

	if len(data.Articles) != 2 {
		t.Errorf("Expected 2 articles, got %d", len(data.Articles))
	}

	for i, tc := range newsTestCases {
		if data.Articles[i].Ticker != tc.value.([]Article)[i].Ticker {
			t.Errorf("Expected ticker to be %s, got %s", tc.value.([]Article)[i].Ticker, data.Articles[i].Ticker)
		}
		if data.Articles[i].Uuid != tc.value.([]Article)[i].Uuid {
			t.Errorf("Expected uuid to be %s, got %s", tc.value.([]Article)[i].Uuid, data.Articles[i].Uuid)
		}
		if data.Articles[i].Title != tc.value.([]Article)[i].Title {
			t.Errorf("Expected title to be %s, got %s", tc.value.([]Article)[i].Title, data.Articles[i].Title)
		}
		if data.Articles[i].Publisher != tc.value.([]Article)[i].Publisher {
			t.Errorf("Expected publisher to be %s, got %s", tc.value.([]Article)[i].Publisher, data.Articles[i].Publisher)
		}
		if data.Articles[i].Link != tc.value.([]Article)[i].Link {
			t.Errorf("Expected link to be %s, got %s", tc.value.([]Article)[i].Link, data.Articles[i].Link)
		}
		if data.Articles[i].ProviderPublishTime != tc.value.([]Article)[i].ProviderPublishTime {
			t.Errorf("Expected provider publish time to be %d, got %d", tc.value.([]Article)[i].ProviderPublishTime, data.Articles[i].ProviderPublishTime)
		}
		if data.Articles[i].Type != tc.value.([]Article)[i].Type {
			t.Errorf("Expected type to be %s, got %s", tc.value.([]Article)[i].Type, data.Articles[i].Type)
		}
		if data.Articles[i].RelatedTickers[0] != tc.value.([]Article)[i].RelatedTickers[0] {
			t.Errorf("Expected related tickers to be %s, got %s", tc.value.([]Article)[i].RelatedTickers[0], data.Articles[i].RelatedTickers[0])
		}
		if data.Articles[i].RelatedTickers[1] != tc.value.([]Article)[i].RelatedTickers[1] {
			t.Errorf("Expected related tickers to be %s, got %s", tc.value.([]Article)[i].RelatedTickers[1], data.Articles[i].RelatedTickers[1])
		}
	}
}

func TestYahooNewsApi(t *testing.T) {
	data, err := NewNews("GOOG")
	if err != nil {
		t.Error("Error fetching data for news", err)
	}

	if len(data.Articles) == 0 {
		t.Error("No articles found")
	}
}
