package plutus

import "strings"

// Helper function to check if the provided attribute belongs to the primary stock on the page
func isPrimary(attr string) bool {
	// Set precendence of pre and post market values over normal values (if pre/post data exists)
	return attr == "" || strings.HasPrefix(attr, "pre") || strings.HasPrefix(attr, "post")
}

// Helper function to clean number strings.
func cleanNumber(s string) string {
	replacer := strings.NewReplacer("%", "", "(", "", ")", "", ",", "")
	return replacer.Replace(s)
}
