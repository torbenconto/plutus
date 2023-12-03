package main

import "fmt"

func main() {
	stock, _ := NewStock("amd")

	fmt.Println(stock.Price)
}
