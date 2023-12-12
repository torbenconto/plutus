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

### Example based documentation
Please use the provided examples to guide you to using plutus to it's full potential.
#### https://github.com/torbenconto/plutus/blob/master/examples


# Future Features
- [-] Related News Articles
- [-] Historical Data
- [-] Price Estimates


And More..