package historical

// Frequency Enum for the frequency of the historical data
type Frequency string

const (
	// Daily Frequency for daily data
	Daily Frequency = "1d"
	// Weekly Frequency for weekly data
	Weekly Frequency = "1wk"
	// Monthly Frequency for monthly data
	Monthly Frequency = "1mo"
)
