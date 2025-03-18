package plutus

import (
	"io"
	"net/http"
)

func getCrumb() (string, error) {
	req, err := http.NewRequest(http.MethodGet, crumbUrl, nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("User-Agent", useragent)
	req.Header.Set("Cookie", cookie)

	get, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer get.Body.Close()

	if get.StatusCode != http.StatusOK {
		return "", err
	}

	body, err := io.ReadAll(get.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
