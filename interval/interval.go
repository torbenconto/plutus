package interval

// Interval Enum for the frequency of the historical data
type Interval string

const (
	// OneMinute Interval for one minute of historical data
	OneMinute Interval = "1m"
	// TwoMinutes Interval for two minutes of historical data
	TwoMinutes Interval = "2m"
	// FiveMinutes Interval for five minutes of historical data
	FiveMinutes Interval = "5m"
	// FifteenMinutes Interval for fifteen minutes of historical data
	FifteenMinutes Interval = "15m"
	// ThirtyMinutes Interval for thirty minutes of historical data
	ThirtyMinutes Interval = "30m"
	// SixtyMinutes Interval for sixty minutes of historical data
	SixtyMinutes Interval = "60m"
	// NinetyMinutes Interval for ninety minutes of historical data
	NinetyMinutes Interval = "90m"
	// OneHour Interval for one hour of historical data
	OneHour Interval = "1h"
	// OneDay Interval for one day of historical data
	OneDay Interval = "1d"
	// FiveDay Interval for five days of historical data
	FiveDay Interval = "5d"
	// OneWeek Interval for one week of historical data
	OneWeek Interval = "1wk"
	// OneMonth Interval for one month of historical data
	OneMonth Interval = "1mo"
	// ThreeMonths Interval for three months of historical data
	ThreeMonths Interval = "3mo"
)
