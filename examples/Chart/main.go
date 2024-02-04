package main

import (
	"fmt"
	"github.com/torbenconto/plutus/historical"
	"github.com/torbenconto/plutus/interval"
	_range "github.com/torbenconto/plutus/range"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"log"
	"time"
)

func main() {
	stock, _ := historical.NewHistorical("GOOG", _range.OneDay, interval.OneMinute)

	data := stock.Data

	// Remove 2nd to last element from data
	//data = append(data[:len(data)-2], data[len(data)-1:]...)

	// Extract x and y values from TimePricePair
	var xs []time.Time
	var ys []float64
	for _, pair := range data {
		xs = append(xs, time.Unix(pair.Time, 0))
		ys = append(ys, pair.Close)
	}

	// Create a new plot
	p := plot.New()

	// Create a new scatter plotter
	pts := make(plotter.XYs, len(xs))
	for i := range pts {
		pts[i].X = float64(xs[i].Unix())
		pts[i].Y = ys[i]
	}

	s, _ := plotter.NewScatter(pts)

	// Add scatter plot to the plot
	p.Add(s)

	// Set plot title and axis labels
	p.Title.Text = "AMD"
	p.X.Label.Text = "Time"
	p.Y.Label.Text = "Price"

	// Format the x-axis ticks to display date and time
	p.X.Tick.Marker = plot.TimeTicks{Format: "2006-01-02\n15:04:05"}

	// Save the plot to a file
	if err := p.Save(800, 600, "price_over_time.png"); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Plot saved as price_over_time.png")
}
