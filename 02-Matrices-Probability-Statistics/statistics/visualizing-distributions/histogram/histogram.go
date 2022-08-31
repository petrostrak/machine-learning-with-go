package main

import (
	"fmt"
	"os"

	"github.com/go-gota/gota/dataframe"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {

	// Open the csv file.
	f, err := os.Open("iris.data")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()

	// Create the dataframe of the csv file.
	df := dataframe.ReadCSV(f)

	// Create a histogram for each of the featured columns in the
	// dataset.
	for _, col := range df.Names() {

		// If the column is one of the feature columns, let's
		// create a histogram of that values.
		if col != "species" {

			// Create a plotter.Values value and fill it with
			// the values from the respective column of the
			// dataframe.
			v := make(plotter.Values, df.Nrow())
			for i, floatVal := range df.Col(col).Float() {
				v[i] = floatVal
			}

			// Make a plot and set its title.
			p := plot.New()
			p.Title.Text = fmt.Sprintf("Histogram of a %s", col)

			// Create a histogram of values drawn from
			// the standard normal.
			h, err := plotter.NewHist(v, 16)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			// Normalize the histogram.
			h.Normalize(1)

			// Add the histogram to the plot.
			p.Add(h)

			// Save the plot to a PNG file.
			err = p.Save(4*vg.Inch, 4*vg.Inch, col+"_hist.png")
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}
	}
}
