package indicators

import (
	"fmt"
	"github.com/shopspring/decimal"
	"github.com/torbenconto/plutus/historical"
)

// SMA is a struct that contains the period and data needed to calculate the Simple Moving Average
type SMA struct {
	Period int
	Data   []historical.PricePoint
}

func NewSMA(period int, data []historical.PricePoint) *SMA {
	return &SMA{
		Period: period,
		Data:   data,
	}
}

func (s *SMA) Calculate() (float64, error) {
	if s.Period > len(s.Data) {
		return 0, fmt.Errorf("period %v is greater than the data length %v", s.Period, len(s.Data))
	}
	var sum decimal.Decimal
	for i := 0; i < s.Period; i++ {
		sum = sum.Add(decimal.NewFromFloat(s.Data[i].Close))
	}

	average := sum.Div(decimal.NewFromInt(int64(s.Period)))
	final, _ := average.Float64()

	return final, nil
}
