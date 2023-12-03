package main

// Map of table names to corresponding Stock Struct values
var YFTableMap = map[string]string{
	"Previous Close": "PrevClose",
	"Open":           "OpenPrice",
	"Bid":            "BidPrice",
	"Ask":            "AskPrice",
	"Day's Range":    "DayRange",
	"52 Week Range":  "FiftyTwoWeekRange",
	"Volume":         "Volume",
}
