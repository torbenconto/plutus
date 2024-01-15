package table

// YFTableMap Map of table names to corresponding Stock Struct values (yahoo finance)
var YFTableMap = map[string]string{
	"Previous Close":           "PrevClose",
	"Open":                     "OpenPrice",
	"Bid":                      "BidPrice",
	"Ask":                      "AskPrice",
	"Volume":                   "Volume",
	"Avg. Volume":              "AvgVolume",
	"Market Cap":               "MarketCap",
	"Beta (5Y Monthly)":        "Beta",
	"PE Ratio (TTM)":           "PE",
	"EPS (TTM)":                "EPS",
	"Earnings Date":            "EarningsDate",
	"Forward Dividend & Yield": "ForwardDividendAndYield",
	"Ex-Dividend Date":         "ExDividendDate",
	"1y Target Est":            "OneYearTargetEst",
}
