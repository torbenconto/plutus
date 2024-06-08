![logo](./assets/logo.webp)
## Plutus
*A stock data centered golang library*

![](https://img.shields.io/github/go-mod/go-version/torbenconto/plutus)
![GitHub Release](https://img.shields.io/github/v/release/torbenconto/plutus)
![GitHub commit activity](https://img.shields.io/github/commit-activity/w/torbenconto/plutus)
[![CodeFactor](https://www.codefactor.io/repository/github/torbenconto/plutus/badge)](https://www.codefactor.io/repository/github/torbenconto/plutus)
[![Codacy Badge](https://app.codacy.com/project/badge/Grade/21f74246fdbc49348075a4dbc156abf3)](https://app.codacy.com/gh/torbenconto/plutus/dashboard?utm_source=gh&utm_medium=referral&utm_content=&utm_campaign=Badge_grade)
![GitHub License](https://img.shields.io/github/license/torbenconto/plutus)
![Lines of code](https://img.shields.io/tokei/lines/github/torbenconto/plutus)![GitHub repo size](https://img.shields.io/github/repo-size/torbenconto/plutus)

### Quick Start
Download the library into your existing golang project
```sh
    go get -u github.com/torbenconto/plutus@latest
```

Create a new stock Quote using the ticker of the stock you want data on.
```go
import "github.com/torbenconto/plutus/stock"
stock, err := stock.NewQuote("AMD")
if err != nil {
	fmt.Printf("An error occured: %e", err)
}
```
Done!, now you can access many different aspects of the stock including price, volume, market cap, and many others!
```go
    fmt.Println(stock.RegularMarketPrice)
```

### Usage

#### Quote
```go
import "github.com/torbenconto/plutus/stock"
stock, err := stock.NewQuote("AMD")
if err != nil {
    fmt.Printf("An error occured: %e", err)
}
```
#### Quote Data Stream
```go
import "github.com/torbenconto/plutus/stock"
stock, err := quote.NewQuote("AMD")
if err != nil {
    fmt.Printf("An error occured: %e", err)
}
// Set delay in Milliseconds
delay := time.Second

// Call stream func using Stock object and a given delay
stream := stock.Stream(delay)

// Get updated data and print out most recent stock price. Runs infinently and returns the newest avalible stock data in the form of a plutus.Stock struct
for {
    data := <-stream
    fmt.Println(data.RegularMarketPrice, data.RegularMarketChangePercent)
}
```

#### Historical Data
```go
import "github.com/torbenconto/plutus/historical"
import _range "github.com/torbenconto/plutus/range
import "github.com/torbenconto/plutus/interval"
// Create a new historical data object using the ticker of the stock you want data on as well as the range and interval of the data.
stock, err := historical.NewHistorical("AMD", _range.FiveDay, interval.OneMin)
if err != nil {
    fmt.Printf("An error occured: %e", err)
}

// Returns a list of all the data points as structs containing the time in unix time and the price of the stock at that time.
for _, data := range stock.Data {
    fmt.Println(data.Time, data.Open, data.Close, data.High, data.Low, data.Volume)
}
```

#### Historical Data Stream
```go
import "github.com/torbenconto/plutus/historical"
import _range "github.com/torbenconto/plutus/range
import "github.com/torbenconto/plutus/interval"
// Create a new historical data object using the ticker of the stock you want data on as well as the range and interval of the data.
stock, err := historical.NewHistorical("AMD", _range.FiveDay, interval.OneMin)
if err != nil {
    fmt.Printf("An error occured: %e", err)
}

// Set delay in Milliseconds
delay := time.Second

// Call stream func using Stock object and a given delay
stream := stock.Stream(delay)

// Get updated data and print out most recent stock price. Runs infinently and returns the newest avalible stock data in the form of a plutus.Stock struct
for {
    data := <-stream
    fmt.Println(data.RegularMarketPrice, data.RegularMarketChangePercent)
}
```

#### Custom API url/request headers
```go
import "github.com/torbenconto/plutus/stock"
import "github.com/torbenconto/plutus/config"

// Create a new stock object with a custom API url and request headers
stock, err := quote.NewQuote("AMD", config.Config{
    Url: "https://example.com",
    UserAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3",
	Cookie: "cookie1=cookie1; cookie2=cookie2",
})

// Create a new historical data object with a custom API url and request headers
stock, err := historical.NewHistorical("AMD", _range.FiveDay, interval.OneMin, config.Config{
    Url: "https://example.com",
    UserAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3",
    Cookie: "cookie1=cookie1; cookie2=cookie2",
})
```

#### Dividend Info
```go
import "github.com/torbenconto/plutus/stock"

info, err := stock.NewDividendInfo("T")
if err != nil {
	fmt.Printf("An error occured: %s\n", err)
}
```

# REST api
## The repo containing the api and information about it is contained here [plutus-api](https://github.com/torbenconto/plutus-api)


### Example based documentation
Please use the provided examples to guide you to using plutus to it's full potential.
#### https://github.com/torbenconto/plutus/blob/master/examples

### Tests
To run the tests for the library, simply run the following command in the root of the project.
```sh
    go test ./...
```

# Future Features
- [x] Historical Data
- [x] Price Estimates
- [ ] Crypto Currency Support

And More..