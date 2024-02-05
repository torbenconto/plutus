package tests

import (
	"fmt"
	"github.com/torbenconto/plutus/indicators"
	"testing"
)

func TestRSI(t *testing.T) {
	// Create a new Relative Strength Index object
	rsi := indicators.NewRSI(3, indicatorData)

	// Calculate the Relative Strength Index
	result, err := rsi.Calculate()
	if err != nil {
		t.Errorf("Error calculating RSI: %v", err)
	}

	// Expected result is 100
	fmt.Println(result)
}
