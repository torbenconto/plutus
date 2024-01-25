package historical

import (
	"time"
)

type Historical struct {
	Ticker               string
	Frequency            Frequency
	StartDate            time.Time
	EndDate              time.Time
	IncludeAdjustedClose bool
}

func NewHistorical(ticker string, frequency Frequency, startDate time.Time, endDate time.Time, includeAdjustedClose bool) (*Historical, error) {
	historical := &Historical{
		Ticker:               ticker,
		Frequency:            frequency,
		StartDate:            startDate,
		EndDate:              endDate,
		IncludeAdjustedClose: includeAdjustedClose,
	}

	return historical.Populate()
}

func (h *Historical) Populate() (*Historical, error) {
	//var err error

	return h, nil
}
