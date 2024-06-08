package stock

// url for the stock data
// Takes 2 arguments: crumb, ticker
const quoteUrl = "https://query2.finance.yahoo.com/v7/finance/quote?formatted=false&crumb=%s&lang=en-US&region=US&symbols=%s&corsDomain=finance.yahoo.com"

// url for dividend data
// Takes 1 argument: ticker
const dividendUrl = "https://api.nasdaq.com/api/quote/%s/dividends?assetclass=stocks"
