package tests

import "github.com/torbenconto/plutus/news"

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
	{"Articles", []news.Article{
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
