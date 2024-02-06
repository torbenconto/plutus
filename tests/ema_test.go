package tests

import (
	"github.com/torbenconto/plutus/indicators"
	"testing"
)

func TestEMA(t *testing.T) {
	// Create a new Exponential Moving Average object
	ema := indicators.NewEMA(5, indicatorData2)

	// Calculate the Exponential Moving Average
	result, err := ema.Calculate()
	if err != nil {
		t.Errorf("Error calculating EMA: %v", err)
	}

	// Expected result is 17
	if result != 17 {
		t.Errorf("Expected 17, got %v", result)
	}
}
