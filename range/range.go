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

func (r Range) String() string {
	return string(r)
}

// RangeFromString Return a Range from a string (default: Max)
func RangeFromString(_range string) Range {
	switch _range {
	case "1d":
		return OneDay
	case "5d":
		return FiveDay
	case "1mo":
		return OneMonth
	case "3mo":
		return ThreeMonth
	case "6mo":
		return SixMonth
	case "1y":
		return OneYear
	case "2y":
		return TwoYear
	case "5y":
		return FiveYear
	case "10y":
		return TenYear
	case "ytd":
		return YearToDate
	case "max":
		return Max
	default:
		return Max
	}
}
