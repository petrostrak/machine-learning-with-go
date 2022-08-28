package main

import (
	"fmt"
	"os"

	"github.com/go-gota/gota/dataframe"
)

func main() {

	// Open the CSV file.
	irisFile, err := os.Open("iris.data")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer irisFile.Close()

	// Create a dataframe from the CSV file.
	// The types of the columns will be inferred.
	irisDF := dataframe.ReadCSV(irisFile)

	// Once we have the data parsed into a dataframe
	// we can filter, subset and select our data easily.

	// Create a filter for the dataframe.
	filter := dataframe.F{
		Colname:    "species",
		Comparator: "==",
		Comparando: "Iris-versicolor",
	}

	// Filter the dataframe to see only the rows where
	// the iris species is "Iris-versicolor"
	versicolorDF := irisDF.Filter(filter)
	if versicolorDF.Err != nil {
		fmt.Println(versicolorDF.Err)
		os.Exit(1)
	}

	// Filter the dataframe again, but only select
	// out the sepal_width and species columns.
	versicolorDF = irisDF.Filter(filter).Select([]string{"sepal_width", "species"})

	// Filter and select the dataframe again, but only display
	// the first three results.
	versicolorDF = irisDF.Filter(filter).Select([]string{"sepal_width", "species"}).Subset([]int{0, 1, 2})

	fmt.Println(versicolorDF)
}
