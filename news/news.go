package news

import (
	"encoding/json"
	"fmt"
	"github.com/torbenconto/plutus"
	"github.com/torbenconto/plutus/config"
	"io"
	"net/http"
)

type response struct {
	News []News `json:"news"`
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

	get, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("Error:", err)
		}
	}(get.Body)

	body, err := io.ReadAll(get.Body)
	if err != nil {
		return nil, err
	}

	var newsResponseData response
	err = json.Unmarshal(body, &newsResponseData)
	if err != nil {
		return nil, err
	}

	if len(newsResponseData.News) == 0 {
		return nil, fmt.Errorf("no news found for %s", n.Ticker)
	}

	n.Articles = newsResponseData.News[0].Articles

	return n, nil
}
