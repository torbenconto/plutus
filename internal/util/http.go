package util

import (
	"fmt"
	"github.com/torbenconto/plutus"
	"github.com/torbenconto/plutus/config"
	"io"
	"net/http"
)

// BuildRequestFromConfig modifies a http.Request from a config.Config. url is the un-formatted url, fallback url is the complete, formatted url.
func BuildRequestFromConfig(conf config.Config, unformattedUrl, fullUrl string) (*http.Request, error) {
	var err error
	var req *http.Request

	if conf.Url != unformattedUrl {
		req, err = http.NewRequest("GET", conf.Url, http.NoBody)
		if err != nil {
			return nil, err
		}
	} else {
		req, err = http.NewRequest("GET", fullUrl, http.NoBody)
		if err != nil {
			return nil, err
		}
	}

	if conf.UserAgent != plutus.UserAgent {
		req.Header.Set("User-Agent", conf.UserAgent)
	} else {
		req.Header.Set("User-Agent", plutus.UserAgent)
	}

	if conf.Cookie != plutus.Cookie {
		req.Header.Set("Cookie", conf.Cookie)
	} else {
		req.Header.Set("Cookie", plutus.Cookie)
	}

	return req, nil
}

func PerformRequest(req *http.Request) ([]byte, error) {
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
	return body, nil
}
