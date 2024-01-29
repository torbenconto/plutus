package indicators

import (
	"github.com/torbenconto/plutus/historical"
)

func AccumulationDistributionLine(data []*historical.PricePoint) []float64 {
	var adl []float64
	var previousAdl float64 = 0

	for _, p := range data {
		moneyFlowMultiplier := ((p.Close - p.Low) - (p.High - p.Close)) / (p.High - p.Low)
		moneyFlowVolume := moneyFlowMultiplier * float64(p.Volume)
		previousAdl += moneyFlowVolume
		adl = append(adl, previousAdl)
	}

	return adl
}
