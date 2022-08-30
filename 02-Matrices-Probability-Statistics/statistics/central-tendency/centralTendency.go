// Central Tendency measures the location of the individual values,
// or where the center of the distribution is located (for example,
// along the preceding linear representation)
//
// Measures of central tendency include the following:
//
// Mean : This is what you might commonly refer to as an average. We
// calculate this by summing all of the numbers in the distribution
// and then dividing by the count of the numbers.
//
// Median : If we sort all of the numbers in our distribution from the
// lowest to highest, this is the number that separates the lowest half
// of the numbers from the highest half of the numbers.
//
// Mode : This is the most frequently occurring value in the distribution.
package main

import (
	"fmt"
	"os"

	"github.com/go-gota/gota/dataframe"
	"github.com/montanaflynn/stats"
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

	// Create a dataframe from the csv file.
	df := dataframe.ReadCSV(f)

	// Get the float values from the "sepal_length" column
	// as we will be looking at the measures for the var.
	sepalLength := df.Col("sepal_length").Float()

	// Calculate the MEAN of the variable
	meanVal := stat.Mean(sepalLength, nil)

	// Calculate the MODE of the variable
	modeVal, modeCount := stat.Mode(sepalLength, nil)

	// Calculate the Median of the variable.
	medianVal, err := stats.Median(sepalLength)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Output the results to standard out.
	fmt.Printf("\nSepal Length Summary Statistics:\n")
	fmt.Printf("Mean value: %0.2f\n", meanVal)
	fmt.Printf("Mode value: %0.2f\n", modeVal)
	fmt.Printf("Mode count: %d\n", int(modeCount))
	fmt.Printf("Median value: %0.2f\n\n", medianVal)
}
