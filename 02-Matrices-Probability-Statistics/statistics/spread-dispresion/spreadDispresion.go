// Measures of spread or dispersion quantify how the values of our
// distribution are spread around the center of our distribution.
// Some of the widely used measures that quantify this are as follows:
//
// Maximum: The highest value of the distribution
//
// Minimum: The lowest value of the distribution
//
// Range: The difference between the maximum value and the minimum value
//
// Variance: This measure is calculated by taking each of the values in the
// distribution, calculating each one's difference from the distribution's
// mean, squaring this difference, adding it to the other squared differences,
// and dividing by the number of values in the distribution.
//
// Standard deviation: The square root of the variance
//
// Quantiles/quartiles: Similar to the median, these measures define cut-off
// points in the distribution where a certain number of lower values are below
// the measure and a certain number of higher values are above the measure
package main

import (
	"fmt"
	"os"

	"github.com/go-gota/gota/dataframe"
	"gonum.org/v1/gonum/floats"
	"gonum.org/v1/gonum/stat"
)

func main() {

	// Open the csv file.
	f, err := os.Open("iris.data")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()

	// Create the dataframe from the csv file.
	df := dataframe.ReadCSV(f)

	// Get the float values from the "sepal_length" column
	// as we will be looking at the measures for this variable.
	sepalLength := df.Col("sepal_length").Float()

	// Calculate the Min of the variable.
	minVal := floats.Min(sepalLength)

	// Calculate the Max of the variable.
	maxVal := floats.Max(sepalLength)

	// Calculate the Median of the variable.
	rangeVal := maxVal - minVal

	// Calculate the variance of the variable.
	varianceVal := stat.Variance(sepalLength, nil)

	// Calculate the standard deviation of the variable.
	stdDevVal := stat.StdDev(sepalLength, nil)

	// Sort the values.
	inds := make([]int, len(sepalLength))
	floats.Argsort(sepalLength, inds)

	// Get the Quantiles.
	quant25 := stat.Quantile(0.25, stat.Empirical, sepalLength, nil)
	quant50 := stat.Quantile(0.50, stat.Empirical, sepalLength, nil)
	quant75 := stat.Quantile(0.75, stat.Empirical, sepalLength, nil)

	// Output the results to standard out.
	fmt.Printf("\nSepal Length Summary Statistics:\n")
	fmt.Printf("Max value: %0.2f\n", maxVal)
	fmt.Printf("Min value: %0.2f\n", minVal)
	fmt.Printf("Range value: %0.2f\n", rangeVal)
	fmt.Printf("Variance value: %0.2f\n", varianceVal)
	fmt.Printf("Std Dev value: %0.2f\n", stdDevVal)
	fmt.Printf("25 Quantile: %0.2f\n", quant25)
	fmt.Printf("50 Quantile: %0.2f\n", quant50)
	fmt.Printf("75 Quantile: %0.2f\n\n", quant75)
}
