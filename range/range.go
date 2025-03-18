package _range

type Range string

const (
	OneDay     Range = "1d"
	FiveDay    Range = "5d"
	OneMonth   Range = "1mo"
	ThreeMonth Range = "3mo"
	SixMonth   Range = "6mo"
	OneYear    Range = "1y"
	TwoYear    Range = "2y"
	FiveYear   Range = "5y"
	TenYear    Range = "10y"
	YearToDate Range = "ytd"
	Max        Range = "max"
)

func (r Range) String() string {
	return string(r)
}

// RangeFromString Returns a Range from a string (default: Max)
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
