package quote

import (
	"encoding/json"
	"fmt"
	"github.com/torbenconto/plutus"
	"github.com/torbenconto/plutus/internal/util"
	"io"
	"net/http"
	"strings"
)

type response struct {
	QuoteResponse struct {
		Result []Quote           `json:"result"`
		Error  map[string]string `json:"error"`
	} `json:"quoteResponse"`
}

type Quote struct {
	Language                          string  `json:"language"`
	Region                            string  `json:"region"`
	QuoteType                         string  `json:"quoteType"`
	TypeDisp                          string  `json:"typeDisp"`
	QuoteSourceName                   string  `json:"quoteSourceName"`
	Triggerable                       bool    `json:"triggerable"`
	CustomPriceAlertConfidence        string  `json:"customPriceAlertConfidence"`
	Currency                          string  `json:"currency"`
	MarketState                       string  `json:"marketState"`
	RegularMarketChangePercent        float64 `json:"regularMarketChangePercent"`
	RegularMarketPrice                float64 `json:"regularMarketPrice"`
	Exchange                          string  `json:"exchange"`
	ShortName                         string  `json:"shortName"`
	LongName                          string  `json:"longName"`
	MessageBoardID                    string  `json:"messageBoardId"`
	ExchangeTimezoneName              string  `json:"exchangeTimezoneName"`
	ExchangeTimezoneShortName         string  `json:"exchangeTimezoneShortName"`
	GmtOffSetMilliseconds             int     `json:"gmtOffSetMilliseconds"`
	Market                            string  `json:"market"`
	EsgPopulated                      bool    `json:"esgPopulated"`
	FirstTradeDateMilliseconds        int64   `json:"firstTradeDateMilliseconds"`
	PriceHint                         int     `json:"priceHint"`
	PostMarketChangePercent           float64 `json:"postMarketChangePercent"`
	PostMarketTime                    int64   `json:"postMarketTime"`
	PostMarketPrice                   float64 `json:"postMarketPrice"`
	PostMarketChange                  float64 `json:"postMarketChange"`
	RegularMarketChange               float64 `json:"regularMarketChange"`
	RegularMarketTime                 int64   `json:"regularMarketTime"`
	RegularMarketDayHigh              float64 `json:"regularMarketDayHigh"`
	RegularMarketDayRange             string  `json:"regularMarketDayRange"`
	RegularMarketDayLow               float64 `json:"regularMarketDayLow"`
	RegularMarketVolume               int64   `json:"regularMarketVolume"`
	RegularMarketPreviousClose        float64 `json:"regularMarketPreviousClose"`
	Bid                               float64 `json:"bid"`
	Ask                               float64 `json:"ask"`
	BidSize                           int     `json:"bidSize"`
	AskSize                           int     `json:"askSize"`
	FullExchangeName                  string  `json:"fullExchangeName"`
	FinancialCurrency                 string  `json:"financialCurrency"`
	RegularMarketOpen                 float64 `json:"regularMarketOpen"`
	AverageDailyVolume3Month          int64   `json:"averageDailyVolume3Month"`
	AverageDailyVolume10Day           int64   `json:"averageDailyVolume10Day"`
	FiftyTwoWeekLowChange             float64 `json:"fiftyTwoWeekLowChange"`
	FiftyTwoWeekLowChangePercent      float64 `json:"fiftyTwoWeekLowChangePercent"`
	FiftyTwoWeekRange                 string  `json:"fiftyTwoWeekRange"`
	FiftyTwoWeekHighChange            float64 `json:"fiftyTwoWeekHighChange"`
	FiftyTwoWeekHighChangePercent     float64 `json:"fiftyTwoWeekHighChangePercent"`
	FiftyTwoWeekLow                   float64 `json:"fiftyTwoWeekLow"`
	FiftyTwoWeekHigh                  float64 `json:"fiftyTwoWeekHigh"`
	FiftyTwoWeekChangePercent         float64 `json:"fiftyTwoWeekChangePercent"`
	EarningsTimestamp                 int64   `json:"earningsTimestamp"`
	EarningsTimestampStart            int64   `json:"earningsTimestampStart"`
	EarningsTimestampEnd              int64   `json:"earningsTimestampEnd"`
	TrailingAnnualDividendRate        float64 `json:"trailingAnnualDividendRate"`
	TrailingPE                        float64 `json:"trailingPE"`
	TrailingAnnualDividendYield       float64 `json:"trailingAnnualDividendYield"`
	EpsTrailingTwelveMonths           float64 `json:"epsTrailingTwelveMonths"`
	EpsForward                        float64 `json:"epsForward"`
	EpsCurrentYear                    float64 `json:"epsCurrentYear"`
	PriceEpsCurrentYear               float64 `json:"priceEpsCurrentYear"`
	SharesOutstanding                 int64   `json:"sharesOutstanding"`
	BookValue                         float64 `json:"bookValue"`
	FiftyDayAverage                   float64 `json:"fiftyDayAverage"`
	FiftyDayAverageChange             float64 `json:"fiftyDayAverageChange"`
	FiftyDayAverageChangePercent      float64 `json:"fiftyDayAverageChangePercent"`
	TwoHundredDayAverage              float64 `json:"twoHundredDayAverage"`
	TwoHundredDayAverageChange        float64 `json:"twoHundredDayAverageChange"`
	TwoHundredDayAverageChangePercent float64 `json:"twoHundredDayAverageChangePercent"`
	MarketCap                         int64   `json:"marketCap"`
	ForwardPE                         float64 `json:"forwardPE"`
	PriceToBook                       float64 `json:"priceToBook"`
	SourceInterval                    int     `json:"sourceInterval"`
	ExchangeDataDelayedBy             int     `json:"exchangeDataDelayedBy"`
	AverageAnalystRating              string  `json:"averageAnalystRating"`
	Tradeable                         bool    `json:"tradeable"`
	CryptoTradeable                   bool    `json:"cryptoTradeable"`
	DisplayName                       string  `json:"displayName"`
	Ticker                            string  `json:"symbol"`
	URL                               string
}

// NewQuote creates a new Quote instance for the given ticker. API url is optional
func NewQuote(ticker string, apiUrl ...string) (*Quote, error) {
	quote := &Quote{
		Ticker: strings.ToUpper(ticker),
	}

	if len(apiUrl) > 0 {
		quote.URL = apiUrl[0]
	} else {
		quote.URL = url
	}

	return quote.Populate()
}

func (q *Quote) Populate() (*Quote, error) {
	var req *http.Request
	var err error

	// Get quote
	if strings.Count(q.URL, "%s") < 2 {
		req, err = http.NewRequest("GET", q.URL, nil)
	} else {
		// Get crumb
		crumb, err := util.GetCrumb()
		if err != nil {
			return nil, fmt.Errorf("error getting crumb: %v", err)
		}
		req, err = http.NewRequest("GET", fmt.Sprintf(q.URL, crumb, q.Ticker), nil)
	}

	req.Header.Set("User-Agent", plutus.UserAgent)
	req.Header.Set("Cookie", plutus.Cookie)

	get, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %v", err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("Error:", err)
		}
	}(get.Body)

	body, err := io.ReadAll(get.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %v", err)
	}

	var quoteResponseData response
	err = json.Unmarshal(body, &quoteResponseData)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %v", err)
	}

	// Check if the Error field is not nil
	if quoteResponseData.QuoteResponse.Error != nil {
		return nil, fmt.Errorf("error returned from API: %s, %s", quoteResponseData.QuoteResponse.Error["code"], quoteResponseData.QuoteResponse.Error["description"])
	}

	// Check if the Result field is empty
	if len(quoteResponseData.QuoteResponse.Result) == 0 {
		return nil, fmt.Errorf("error returned from API: no result returned")
	}

	// Url for complete struct data, necessary for Stream method
	quoteResponseData.QuoteResponse.Result[0].URL = q.URL

	return &quoteResponseData.QuoteResponse.Result[0], nil
}
