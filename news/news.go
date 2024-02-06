package news

import "github.com/torbenconto/plutus/config"

type response struct {
	News []News `json:"news"`
}

type News struct {
	Ticker              string   `json:"ticker"`
	Uuid                string   `json:"uuid"`
	Title               string   `json:"title"`
	Publisher           string   `json:"publisher"`
	Link                string   `json:"link"`
	ProviderPublishTime int64    `json:"providerPublishTime"`
	Type                string   `json:"type"`
	RelatedTickers      []string `json:"relatedTickers"`
	Config              config.Config
}

func NewNews(symbol string, config config.Config) (*News, error) {
	return &News{
		Config: config,
	}, nil
}
