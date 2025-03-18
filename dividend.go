package plutus

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type dividendResponse struct {
	Data struct {
		DividendHeaderValues []struct {
			Label string
			Value string
		} `json:"dividendHeaderValues"`
	} `json:"data"`
	Status struct {
		ResponseCode int `json:"rCode"`
	} `json:"status"`
}

type DividendInfo struct {
	ExDividendDate       time.Time
	DividendYield        float64
	AnnualDividendAmount float64
	Ticker               string
}

func GetDividendInfo(symbol string, url ...string) (*DividendInfo, error) {
	var usedUrl string
	if len(url) > 0 {
		usedUrl = url[0]
	} else {
		usedUrl = dividendUrl(symbol)
	}

	req, err := http.NewRequest(http.MethodGet, usedUrl, nil)
	if err != nil {
		return nil, fmt.Errorf("could not create request: %v", err)
	}

	req.Header.Set("User-Agent", useragent)
	req.Header.Set("Cookie", cookie)
	req.Header.Set("Accept", "application/json")

	get, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("could not send request: %v", err)
	}
	defer get.Body.Close()

	if get.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("could not get quote: %v", err)
	}

	body, err := io.ReadAll(get.Body)
	if err != nil {
		return nil, fmt.Errorf("could not read response: %v", err)
	}

	var dr dividendResponse
	if err := json.Unmarshal(body, &dr); err != nil {
		return nil, fmt.Errorf("could not unmarshal response: %v", err)
	}

	if dr.Status.ResponseCode != 200 {
		return nil, fmt.Errorf("error returned from API: %d", dr.Status.ResponseCode)
	}

	if len(dr.Data.DividendHeaderValues) == 0 {
		return nil, fmt.Errorf("error returned from API: no result returned")
	}

	var d DividendInfo

	for _, value := range dr.Data.DividendHeaderValues {
		if value.Value == "N/A" {
			return nil, fmt.Errorf("error: dividend data contains N/A for %s, the stock may not pay a dividend or the api has not yet updated", value.Label)
		}

		switch value.Label {
		case "Ex-Dividend Date":
			date, err := time.Parse("1/02/2006", value.Value)
			if err != nil {
				return nil, fmt.Errorf("error parsing date: %v", err)
			}
			d.ExDividendDate = date

		case "Dividend Yield":
			yield, err := strconv.ParseFloat(strings.TrimSuffix(value.Value, "%"), 64)
			if err != nil {
				return nil, fmt.Errorf("error parsing into struct: %v", err)
			}
			d.DividendYield = yield

		case "Annual Dividend":
			dividend, err := strconv.ParseFloat(strings.TrimPrefix(value.Value, "$"), 64)
			if err != nil {
				return nil, fmt.Errorf("error parsing into struct: %v", err)
			}
			d.AnnualDividendAmount = dividend
		}
	}

	d.Ticker = symbol

	return &d, nil
}
