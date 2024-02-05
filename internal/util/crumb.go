package util

import (
	"fmt"
	"github.com/torbenconto/plutus"
	"io"
	"net/http"
)

func GetCrumb() (string, error) {

	req, err := http.NewRequest("GET", "https://query1.finance.yahoo.com/v1/test/getcrumb", nil)
	if err != nil {
		return "", fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("User-Agent", plutus.UserAgent)
	req.Header.Set("Cookie", plutus.Cookie)

	get, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("error sending request: %v", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("Error:", err)
		}
	}(get.Body)

	body, err := io.ReadAll(get.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response: %v", err)
	}

	return string(body), nil
}
