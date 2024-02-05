package tests

import (
	"fmt"
	"github.com/torbenconto/plutus/historical"
	"github.com/torbenconto/plutus/indicators"
	"testing"
)

func TestSMA(t *testing.T) {
	data := []historical.PricePoint{
		{
			Time:   1,
			Open:   1.0,
			Close:  2.7,
			High:   3.2,
			Low:    0.9,
			Volume: 20100,
		},
		{
			Time:   2,
			Open:   2.0,
			Close:  3.7,
			High:   4.2,
			Low:    1.9,
			Volume: 20200,
		},
		{
			Time:   3,
			Open:   3.0,
			Close:  4.7,
			High:   5.2,
			Low:    2.9,
			Volume: 20300,
		},
	}

	// Create a new Simple Moving Average object
	sma := indicators.NewSMA(3, data)

	// Calculate the Simple Moving Average
	result, err := sma.Calculate()
	if err != nil {
		t.Errorf("Error calculating SMA: %v", err)
	}

	// Expected result is 3.7
	fmt.Println(result)
}
