package interval

type Interval string

const (
	OneMinute      Interval = "1m"
	TwoMinutes     Interval = "2m"
	FiveMinutes    Interval = "5m"
	FifteenMinutes Interval = "15m"
	ThirtyMinutes  Interval = "30m"
	SixtyMinutes   Interval = "60m"
	NinetyMinutes  Interval = "90m"
	OneHour        Interval = "1h"
	OneDay         Interval = "1d"
	FiveDay        Interval = "5d"
	OneWeek        Interval = "1wk"
	OneMonth       Interval = "1mo"
	ThreeMonths    Interval = "3mo"
)

func (i Interval) String() string {
	return string(i)
}

// IntervalFromString Returns an Interval from a string (default: OneDay)
func IntervalFromString(interval string) Interval {
	switch interval {
	case "1m":
		return OneMinute
	case "2m":
		return TwoMinutes
	case "5m":
		return FiveMinutes
	case "15m":
		return FifteenMinutes
	case "30m":
		return ThirtyMinutes
	case "60m":
		return SixtyMinutes
	case "90m":
		return NinetyMinutes
	case "1h":
		return OneHour
	case "1d":
		return OneDay
	case "5d":
		return FiveDay
	case "1wk":
		return OneWeek
	case "1mo":
		return OneMonth
	case "3mo":
		return ThreeMonths
	default:
		return OneDay
	}
}
