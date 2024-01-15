package plutus

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

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

// Helper function to set the struct field based on its type.
func (s *Stock) setField(fieldName string, value string) {
	val := reflect.ValueOf(s).Elem()
	field := val.FieldByName(fieldName)

	value = cleanNumber(value)

	switch field.Kind() {
	case reflect.String:
		field.SetString(value)
	case reflect.Float64:
		fmt.Println(fieldName, value)
		fieldFloat, _ := strconv.ParseFloat(value, 64)
		field.SetFloat(fieldFloat)
	case reflect.Int:
		fieldInt, _ := strconv.Atoi(value)
		field.SetInt(int64(fieldInt))
	}
}
