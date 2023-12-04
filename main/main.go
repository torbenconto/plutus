package main

import (
	"fmt"

	"github.com/torbenconto/plutus"
)

func main() {
	stock, _ := plutus.NewStock("AMD")
	plutus.Stream(stock, 50)

	delayMs := 1000 // Set the delay to 1000 milliseconds (1 second)
	stream := plutus.Stream(stock, delayMs)

	// Print the first 5 elements of the stream
	for i := 0; i < 5; i++ {
		data := <-stream
		fmt.Println(data)
	}
}
