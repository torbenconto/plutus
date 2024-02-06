package tests

import "github.com/torbenconto/plutus/historical"

var indicatorData = []historical.PricePoint{
	{
		Time:   1,
		Open:   1.0,
		Close:  2.7,
		High:   3.2,
		Low:    0.9,
		Volume: 20100,
	},
	{
		Time:   2,
		Open:   2.0,
		Close:  3.7,
		High:   4.2,
		Low:    1.9,
		Volume: 20200,
	},
	{
		Time:   3,
		Open:   3.0,
		Close:  4.7,
		High:   5.2,
		Low:    2.9,
		Volume: 20300,
	},
}

var indicatorData2 = []historical.PricePoint{
	{Close: 10.0},
	{Close: 11.0},
	{Close: 12.0},
	{Close: 13.0},
	{Close: 14.0},
	{Close: 15.0},
	{Close: 16.0},
	{Close: 17.0},
	{Close: 18.0},
	{Close: 19.0},
}
