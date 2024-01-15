package util

import "strings"

// IsPrimary Helper function to check if the provided attribute belongs to the primary quote on the page
func IsPrimary(attr string) bool {
	// Set precendence of pre and post market values over normal values (if pre/post data exists)
	return attr == "" || strings.HasPrefix(attr, "pre") || strings.HasPrefix(attr, "post")
}

// CleanNumber Helper function to clean number strings.
func CleanNumber(s string) string {
	replacer := strings.NewReplacer("%", "", "(", "", ")", "", ",", "")
	return replacer.Replace(s)
}
