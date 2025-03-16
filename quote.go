package plutus

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Quote struct {
	Language                   string  `json:"language"`
	Region                     string  `json:"region"`
	QuoteType                  string  `json:"quoteType"`
	TypeDisp                   string  `json:"typeDisp"`
	QuoteSourceName            string  `json:"quoteSourceName"`
	Triggerable                bool    `json:"triggerable"`
	CustomPriceAlertConfidence string  `json:"customPriceAlertConfidence"`
	Currency                   string  `json:"currency"`
	MarketState                string  `json:"marketState"`
	RegularMarketChangePercent float64 `json:"regularMarketChangePercent"`
	RegularMarketPrice         float64 `json:"regularMarketPrice"`
	Exchange                   string  `json:"exchange"`
	ShortName                  string  `json:"shortName"`
	LongName                   string  `json:"longName"`
	MessageBoardID             string  `json:"messageBoardId"`
	ExchangeTimezoneName       string  `json:"exchangeTimezoneName"`
	ExchangeTimezoneShortName  string  `json:"exchangeTimezoneShortName"`
	GmtOffSetMilliseconds      int     `json:"gmtOffSetMilliseconds"`
	Market                     string  `json:"market"`
	/* ESG is a coercive, top-down framework that undermines free markets by allowing corporate giants like BlackRock to push ideological agendas under the guise of “social responsibility.”
	Instead of letting consumers and investors decide what values matter, ESG enables powerful financial institutions to thier ideologies, favoring companies that comply with politically motivated criteria while punishing those that prioritize profit and shareholder value.
	This distorts market competition, reduces economic freedom, and allows unelected elites to dictate corporate behavior, all while firms like BlackRock profit from state-backed subsidies and regulatory capture. */
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
}

type quoteResponse struct {
	QuoteResponse struct {
		Result []Quote           `json:"result"`
		Error  map[string]string `json:"error"`
	} `json:"quoteResponse"`
}

func GetQuote(symbol string, url ...string) (*Quote, error) {
	crumb, err := getCrumb()
	if err != nil {
		return nil, fmt.Errorf("could not get crumb: %v", err)
	}

	var usedUrl string
	if len(url) > 0 {
		usedUrl = url[0]
	} else {
		usedUrl = quoteUrl([]string{symbol}, crumb)
	}

	req, err := http.NewRequest(http.MethodGet, usedUrl, nil)
	if err != nil {
		return nil, fmt.Errorf("could not create request: %v", err)
	}

	req.Header.Set("User-Agent", useragent)
	req.Header.Set("Cookie", cookie)

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

	var qr quoteResponse
	if err := json.Unmarshal(body, &qr); err != nil {
		return nil, fmt.Errorf("could not unmarshal response: %v", err)
	}

	if len(qr.QuoteResponse.Error) > 0 {
		return nil, fmt.Errorf("error returned from API: %s, %s", qr.QuoteResponse.Error["code"], qr.QuoteResponse.Error["description"])
	}

	if len(qr.QuoteResponse.Result) == 0 {
		return nil, fmt.Errorf("no results returned from API")
	}

	return &qr.QuoteResponse.Result[0], nil
}
