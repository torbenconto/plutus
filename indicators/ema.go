package indicators

import (
	"fmt"
	"github.com/shopspring/decimal"
	"github.com/torbenconto/plutus/historical"
)

type EMA struct {
	Period int
	Data   []historical.PricePoint
}

func NewEMA(period int, data []historical.PricePoint) *EMA {
	return &EMA{
		Period: period,
		Data:   data,
	}
}

func (e *EMA) Calculate() (float64, error) {
	if e.Period > len(e.Data) {
		return 0, fmt.Errorf("period %v is greater than the data length %v", e.Period, len(e.Data))
	}

	// Calculate initial SMA
	sma := NewSMA(e.Period, e.Data)
	smaResult, err := sma.Calculate()
	if err != nil {
		return 0, fmt.Errorf("error calculating SMA: %v", err)
	}

	smaResultDecimal := decimal.NewFromFloat(smaResult)

	multiplier := decimal.NewFromFloat(2.0).Div(decimal.NewFromFloat(float64(e.Period) + 1))

	ema := smaResultDecimal
	for i := e.Period; i < len(e.Data); i++ {
		closePriceDecimal := decimal.NewFromFloat(e.Data[i].Close)
		ema = closePriceDecimal.Sub(ema).Mul(multiplier).Add(ema)
	}

	emaFloat64, _ := ema.Float64()
	return emaFloat64, nil
}
