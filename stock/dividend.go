package stock

import (
	"encoding/json"
	"fmt"
	"github.com/torbenconto/plutus"
	"github.com/torbenconto/plutus/config"
	"github.com/torbenconto/plutus/internal/util"
	"log"
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
	Config               config.Config
}

func NewDividendInfo(ticker string, dividendConfig ...config.Config) (*DividendInfo, error) {
	dividendInfo := &DividendInfo{
		Ticker: strings.ToUpper(ticker),
	}

	if len(dividendConfig) > 0 {
		dividendInfo.Config = dividendConfig[0]
	} else {
		dividendInfo.Config = config.Config{
			Url:       dividendUrl,
			UserAgent: plutus.UserAgent,
			Cookie:    plutus.Cookie,
		}
	}

	return dividendInfo.Populate()
}

func (d *DividendInfo) Populate() (*DividendInfo, error) {
	var req *http.Request
	var err error

	req, err = util.BuildRequestFromConfig(req, d.Config, dividendUrl, fmt.Sprintf(dividendUrl, d.Ticker))
	if err != nil {
		return nil, fmt.Errorf("error building request %v", err)
	}

	body, err := util.PerformRequest(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request %v", err)
	}

	var dividendResponseData dividendResponse
	err = json.Unmarshal(body, &dividendResponseData)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %v", err)
	}

	if dividendResponseData.Status.ResponseCode != 200 {
		// No need to return the actual error code because the nasdaq api always returns the same error message "Something went wrong.Please try again later."
		return nil, fmt.Errorf("error returned from API: non 200 status code")
	}

	if len(dividendResponseData.Data.DividendHeaderValues) == 0 {
		return nil, fmt.Errorf("error returned from API: no result returned")
	}

	// No streaming of dividend info so no need to complete struct with config

	for _, value := range dividendResponseData.Data.DividendHeaderValues {
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
				log.Fatal(err)
			}
			d.DividendYield = yield

		case "Annual Dividend":
			dividend, err := strconv.ParseFloat(strings.TrimPrefix(value.Value, "$"), 64)
			if err != nil {
				log.Fatal(err)
			}
			d.AnnualDividendAmount = dividend
		}
	}

	return d, nil
}
