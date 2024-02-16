package news

import (
	"encoding/json"
	"fmt"
	"github.com/torbenconto/plutus"
	"github.com/torbenconto/plutus/config"
	"github.com/torbenconto/plutus/internal/util"
	"net/http"
)

type response struct {
	News []Article `json:"news"`
}

type Article struct {
	Ticker              string   `json:"ticker"`
	Uuid                string   `json:"uuid"`
	Title               string   `json:"title"`
	Publisher           string   `json:"publisher"`
	Link                string   `json:"link"`
	ProviderPublishTime int64    `json:"providerPublishTime"`
	Type                string   `json:"type"`
	RelatedTickers      []string `json:"relatedTickers"`
}

type News struct {
	Ticker   string
	Articles []Article
	Config   config.Config
}

func NewNews(ticker string, newsConfig ...config.Config) (*News, error) {
	news := &News{
		Ticker: ticker,
	}

	if len(newsConfig) > 0 {
		news.Config = newsConfig[0]
	} else {
		news.Config = config.Config{
			Url:       url,
			UserAgent: plutus.UserAgent,
			Cookie:    plutus.Cookie,
		}
	}

	return news.Populate()
}

func (n *News) Populate() (*News, error) {
	var req *http.Request
	var err error

	req, err = util.BuildRequestFromConfig(req, n.Config, url, fmt.Sprintf(url, n.Ticker))
	if err != nil {
		return nil, fmt.Errorf("error building request: %v", err)
	}

	body, err := util.MakeRequest(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %v", err)
	}

	var newsResponseData response
	err = json.Unmarshal(body, &newsResponseData)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %v", err)
	}

	if len(newsResponseData.News) == 0 {
		return nil, fmt.Errorf("no news found for %s", n.Ticker)
	}

	n.Articles = newsResponseData.News

	return n, nil
}
