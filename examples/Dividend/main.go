package main

import (
	"fmt"
	"github.com/torbenconto/plutus/stock"
)

func main() {
	// Create new DividendInfo struct which will be auto populated
	info, err := stock.NewDividendInfo("T")
	if err != nil {
		fmt.Printf("An error occured: %s\n", err)
	}

	// Use dividend data
	fmt.Printf("Yield: %.2f\n", info.DividendYield)
	fmt.Printf("Annual Amount: %.2f\n", info.AnnualDividendAmount)
	fmt.Printf("Ex Date: %s", info.ExDividendDate)
}
