![logo](./assets/logo.webp)
## Plutus
*A stock data centered golang library*

### Quick Start
Download the library into your existing golang project
```sh
    go get -u github.com/torbenconto/plutus@latest
```

Create a new stock object using the ticker of the stock you want data on and the website you want data from (in this case yahoo finance is used)
```go
stock, err := plutus.NewStock("AMD", plutus.YahooFinanceProvider)
if err != nil {
	fmt.Printf("An error occured: %e", err)
}
```
Done!, now you can access many different aspects of the stock including price, volume, market cap, and many others!
```go
    fmt.Println(stock.Price)
```

### Providers
Plutus currently supports 2 different providers for stock data, Yahoo Finance and Alpha Vantage.
#### Yahoo Finance (Web Scraping)
Yahoo Finance is the default provider for plutus and is the most reliable provider. It is free to use as it scrapes data, provides the most data.
```go
    stock, err := plutus.NewStock("ticker", plutus.YahooFinanceProvider)
```
#### Alpha Vantage (Api Key Required)
Alpha Vantage is the fastest provider for plutus and is the most reliable provider. It is free to use at a certain tier but not reccomended to use without a premium subscription as it doesen't provide realtime data and limits your requests. Currently provides less data that yfinance but is being worked on to provide more data.
```go
    stock, err := plutus.NewStock("ticker", plutus.AlphaVantageProvider, "api_key_here")
```

### Example based documentation
Please use the provided examples to guide you to using plutus to it's full potential.
#### https://github.com/torbenconto/plutus/blob/master/examples


# Future Features
- [-] Related News Articles
- [-] Historical Data
- [-] Price Estimates


And More..