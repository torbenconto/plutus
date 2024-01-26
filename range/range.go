package _range

// Range Enum for the range of the historical data
type Range string

const (
	// OneDay Range for one day of historical data
	OneDay Range = "1d"
	// FiveDay Range for five days of historical data
	FiveDay Range = "5d"
	// OneMonth Range for one month of historical data
	OneMonth Range = "1mo"
	// ThreeMonth Range for three months of historical data
	ThreeMonth Range = "3mo"
	// SixMonth Range for six months of historical data
	SixMonth Range = "6mo"
	// OneYear Range for one year of historical data
	OneYear Range = "1y"
	// TwoYear Range for two years of historical data
	TwoYear Range = "2y"
	// FiveYear Range for five years of historical data
	FiveYear Range = "5y"
	// TenYear Range for ten years of historical data
	TenYear Range = "10y"
	// YearToDate Range for year to date of historical data
	YearToDate Range = "ytd"
	// Max Range for all historical data
	Max Range = "max"
)
