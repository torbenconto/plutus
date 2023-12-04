![logo](./assets/DALL·E%202023-12-03%2020.30.35%20-%20a%2019th%20century%20greek%20oil%20painting.webp)
## Plutus
*A stock data centered golang library*

### Quick Start
Download the library into your existing golang project
```sh
    go get -u github.com/torbenconto/plutus@latest
```

Create a new stock object using the ticker of the stock you want data on
```go
    stock, err := plutus.NewStock("AMD")
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