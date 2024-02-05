package indicators

import (
	"github.com/shopspring/decimal"
	"github.com/torbenconto/plutus/historical"
)

// RSI is a struct that contains the period and data needed to calculate the Relative Strength Index
type RSI struct {
	Period int
	Data   []historical.PricePoint
}

func NewRSI(period int, data []historical.PricePoint) *RSI {
	return &RSI{
		Period: period,
		Data:   data,
	}
}

func (r *RSI) Calculate() (float64, error) {
	if r.Period > len(r.Data) {
		return 0, nil
	}

	gain := decimal.NewFromFloat(0)
	loss := decimal.NewFromFloat(0)

	for i := 1; i < r.Period; i++ {
		change := decimal.NewFromFloat(r.Data[i].Close).Sub(decimal.NewFromFloat(r.Data[i-1].Close))
		if change.GreaterThan(decimal.NewFromFloat(0)) {
			gain = gain.Add(change)
		} else {
			loss = loss.Add(change.Neg())
		}
	}

	averageGain := gain.Div(decimal.NewFromFloat(float64(r.Period)))
	averageLoss := loss.Div(decimal.NewFromFloat(float64(r.Period)))

	rs := averageGain.Div(averageLoss)
	rsAsFloat, _ := rs.Float64()
	rsi := 100 - (100 / (1 + rsAsFloat))
	return rsi, nil
}
