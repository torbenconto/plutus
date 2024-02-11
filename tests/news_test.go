package tests

import (
	"github.com/torbenconto/plutus/config"
	"github.com/torbenconto/plutus/news"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewsPopulate(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, err := w.Write(newsServerData)
		if err != nil {
			t.Error("Error writing response")
		}
	}))
	defer server.Close()

	data, err := news.NewNews("GOOG", config.Config{
		Url: server.URL,
	})
	if err != nil {
		t.Error("Error fetching data for news", err)
	}

	if len(data.Articles) != 2 {
		t.Errorf("Expected 2 articles, got %d", len(data.Articles))
	}

	for i, tc := range newsTestCases {
		if data.Articles[i].Ticker != tc.value.([]news.Article)[i].Ticker {
			t.Errorf("Expected ticker to be %s, got %s", tc.value.([]news.Article)[i].Ticker, data.Articles[i].Ticker)
		}
		if data.Articles[i].Uuid != tc.value.([]news.Article)[i].Uuid {
			t.Errorf("Expected uuid to be %s, got %s", tc.value.([]news.Article)[i].Uuid, data.Articles[i].Uuid)
		}
		if data.Articles[i].Title != tc.value.([]news.Article)[i].Title {
			t.Errorf("Expected title to be %s, got %s", tc.value.([]news.Article)[i].Title, data.Articles[i].Title)
		}
		if data.Articles[i].Publisher != tc.value.([]news.Article)[i].Publisher {
			t.Errorf("Expected publisher to be %s, got %s", tc.value.([]news.Article)[i].Publisher, data.Articles[i].Publisher)
		}
		if data.Articles[i].Link != tc.value.([]news.Article)[i].Link {
			t.Errorf("Expected link to be %s, got %s", tc.value.([]news.Article)[i].Link, data.Articles[i].Link)
		}
		if data.Articles[i].ProviderPublishTime != tc.value.([]news.Article)[i].ProviderPublishTime {
			t.Errorf("Expected provider publish time to be %d, got %d", tc.value.([]news.Article)[i].ProviderPublishTime, data.Articles[i].ProviderPublishTime)
		}
		if data.Articles[i].Type != tc.value.([]news.Article)[i].Type {
			t.Errorf("Expected type to be %s, got %s", tc.value.([]news.Article)[i].Type, data.Articles[i].Type)
		}
		if data.Articles[i].RelatedTickers[0] != tc.value.([]news.Article)[i].RelatedTickers[0] {
			t.Errorf("Expected related tickers to be %s, got %s", tc.value.([]news.Article)[i].RelatedTickers[0], data.Articles[i].RelatedTickers[0])
		}
		if data.Articles[i].RelatedTickers[1] != tc.value.([]news.Article)[i].RelatedTickers[1] {
			t.Errorf("Expected related tickers to be %s, got %s", tc.value.([]news.Article)[i].RelatedTickers[1], data.Articles[i].RelatedTickers[1])
		}
	}
}

func TestYahooNewsApi(t *testing.T) {
	data, err := news.NewNews("GOOG")
	if err != nil {
		t.Error("Error fetching data for news", err)
	}

	if len(data.Articles) == 0 {
		t.Error("No articles found")
	}
}
