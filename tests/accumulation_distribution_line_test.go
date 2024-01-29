package tests

import (
	"github.com/torbenconto/plutus/historical"
	"github.com/torbenconto/plutus/indicators"
	"testing"
)

func TestAccumulationDistributionLine(t *testing.T) {
	data := []*historical.PricePoint{
		{High: 10, Low: 5, Close: 7, Volume: 1000},
		{High: 15, Low: 10, Close: 12, Volume: 2000},
		{High: 20, Low: 15, Close: 17, Volume: 3000},
	}

	expected := []float64{-200, -600, -1200}

	result := indicators.AccumulationDistributionLine(data)

	for i, v := range result {
		if v != expected[i] {
			t.Errorf("Expected %f but got %f at index %d", expected[i], v, i)
		}
	}
}
