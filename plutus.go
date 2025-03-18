package plutus

import (
	"fmt"
	"strings"
)

const crumbUrl = "https://query1.finance.yahoo.com/v1/test/getcrumb"

func quoteUrl(symbols []string, crumb string) string {
	return fmt.Sprintf("https://query2.finance.yahoo.com/v7/finance/quote?formatted=false&symbols=%s&lang=en-US&region=US&crumb=%s&corsDomain=finance.yahoo.com", strings.Join(symbols, ","), crumb)
}

func historicalUrl(symbol string, _range string, interval string) string {
	return fmt.Sprintf("https://query1.finance.yahoo.com/v8/finance/chart/%s?region=US&lang=en-US&includePrePost=false&range=%s&interval=%s&useYfid=true&corsDomain=finance.yahoo.com&.tsrc=finance&indicators=quote", symbol, _range, interval)
}

const useragent = "Mozilla/5.0 Plutus/2.0"
const cookie = "A1=d=AQABBJa6ZGUCENTrtXZe-SIkJXEfX8ySTV8FEgEACAINl2XBZdwx0iMA_eMBAAcIlrpkZcySTV8ID4OikZ64YJY1AyPq3hVnIwkBBwoBYg&S=AQAAAgfHJ22fM3FOY15li56pu8k;"

func dividendUrl(symbol string) string {
	return fmt.Sprintf("https://api.nasdaq.com/api/quote/%s/dividends?assetclass=stocks", symbol)
}
