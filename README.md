![logo](./assets/logo.webp)
## Plutus
*A stock data centered golang library*

### Quick Start
Download the library into your existing golang project
```sh
    go get -u github.com/torbenconto/plutus@latest
```

Create a new stock Quote using the ticker of the stock you want data on.
```go
import "github.com/torbenconto/plutus/quote"
stock, err := quote.NewQuote("AMD")
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
import "github.com/torbenconto/plutus/quote"
stock, err := quote.NewQuote("AMD")
if err != nil {
    fmt.Printf("An error occured: %e", err)
}
```
#### Quote Data Stream
```go
import "github.com/torbenconto/plutus/quote"
stock, err := quote.NewQuote("AMD")
if err != nil {
    fmt.Printf("An error occured: %e", err)
}
// Set delay in Milliseconds
delayInMS := 1000

// Call stream func using Stock object and a given delay
stream := stock.Stream(delayInMS)

// Get updated data and print out most recent quote price. Runs infinently and returns the newest avalible quote data in the form of a plutus.Stock struct
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



# REST api
## The repo containing the api and information about it is contained here [plutus-api](https://github.com/torbenconto/plutus-api)


### Example based documentation
Please use the provided examples to guide you to using plutus to it's full potential.
#### https://github.com/torbenconto/plutus/blob/master/examples


# Future Features
- [ ] Related News Articles
- [x] Historical Data
- [x] Price Estimates
- [ ] Crypto Currency Support


And More..