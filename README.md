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
    fmt.Println(stock.Price)
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
    fmt.Println(data.Price, data.ChangePercent)
}
```


# Api Documentation


### Example based documentation
Please use the provided examples to guide you to using plutus to it's full potential.
#### https://github.com/torbenconto/plutus/blob/master/examples


# Future Features
- [-] Related News Articles
- [-] Historical Data
- [-] Price Estimates


And More..