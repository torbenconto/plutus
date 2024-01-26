package main

import (
	"fmt"
	"github.com/torbenconto/plutus/historical"
	"github.com/torbenconto/plutus/interval"
	_range "github.com/torbenconto/plutus/range"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"log"
)

func main() {
	stock, _ := historical.NewHistorical("AMD", _range.OneDay, interval.OneHour)

	data := stock.Data

	// Extract x and y values from TimePricePair
	var xs []float64
	var ys []float64
	for _, pair := range data {
		xs = append(xs, float64(pair.Time))
		ys = append(ys, pair.Close)
	}

	// Create a new plot
	p := plot.New()

	// Create a new scatter plotter
	pts := make(plotter.XYs, len(xs))
	for i := range pts {
		pts[i].X = xs[i]
		pts[i].Y = ys[i]
	}

	s, _ := plotter.NewScatter(pts)

	// Add scatter plot to the plot
	p.Add(s)

	// Set plot title and axis labels
	p.Title.Text = "Price over Time"
	p.X.Label.Text = "Time"
	p.Y.Label.Text = "Price"

	// Save the plot to a file
	if err := p.Save(800, 400, "price_over_time.png"); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Plot saved as price_over_time.png")
}
