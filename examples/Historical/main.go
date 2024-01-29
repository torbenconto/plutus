package main

import (
	"fmt"
	"github.com/torbenconto/plutus/historical"
	"github.com/torbenconto/plutus/interval"
	_range "github.com/torbenconto/plutus/range"
)

func main() {
	data, err := historical.NewHistorical("AMD", _range.OneDay, interval.OneMinute)
	if err != nil {
		panic(err)
	}

	// Do something with data
	for _, point := range data.Data {
		fmt.Println(point.High, point.Low, point.Open, point.Close, point.Volume, point.Time)
	}
}
