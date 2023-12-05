package plutus

import "math"

// Helper function to check if the provided attribute belongs to the primary stock on the page
func isPrimary(attr string) bool {
	return attr == ""
}

func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}
