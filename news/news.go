package news

import (
	"fmt"
	"github.com/torbenconto/plutus/config"
	"net/http"
)

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

func (n *News) Populate() (*News, error) {
	var req *http.Request
	var err error

	if n.Config.Url != "" {
		req, err = http.NewRequest("GET", n.Config.Url, nil)
		if err != nil {
			return nil, err
		}
	} else {
		req, err = http.NewRequest("GET", fmt.Sprintf(url, n.Ticker), nil)
		if err != nil {
			return nil, err
		}
	}

	req.Header.Set("User-Agent", n.Config.UserAgent)
	req.Header.Set("Cookie", n.Config.Cookie)

	return n, nil
}
