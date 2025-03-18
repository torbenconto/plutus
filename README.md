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
    go get -u github.com/torbenconto/plutus/v2
```

## Basic Usage
```go
import "github.com/torbenconto/plutus"
stock, err := plutus.GetQuote("AMD")
if err != nil {
	fmt.Printf("An error occured: %e", err)
}
fmt.Printf("The current price of AMD is: %f", stock.RegularMarketPrice)
```

### Historical Data

```go
import "github.com/torbenconto/plutus"
import _range "github.com/torbenconto/plutus/range"
import "github.com/torbenconto/plutus/interval"

// Create a new historical data object using the ticker of the stock you want data on as well as the range and interval of the data.
stock, err := plutus.GetHistoricalQuote("AMD", _range.FiveDay, interval.OneMin)
if err != nil {
    fmt.Printf("An error occured: %e", err)
}

// Returns a list of all the data points as structs containing the time in unix time and the price of the stock at that time.
for _, data := range stock.Data {
    fmt.Println(data.Time, data.Open, data.Close, data.High, data.Low, data.Volume)
}
```

### Dividends
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

# Notes
Please only use this module for personal use.
Don't overload my deployment or ill get really sad 🥲.