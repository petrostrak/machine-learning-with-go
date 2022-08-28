package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func main() {

	// Open the iris dataset file.
	f, err := os.Open("iris.data")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()

	// Create a new CSV reader reading from
	// the opened file.
	r := csv.NewReader(f)

	// Assume we don't know the number of fields
	// per line. By setting FieldsPerRecord negative,
	// each row may have a variable number of fields.
	r.FieldsPerRecord = -1

	// rawCSVData will hold our successfully parsed rows.
	var rawCSVData [][]string

	for {
		// Read in a row
		record, err := r.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Append the record to our dataset.
		rawCSVData = append(rawCSVData, record)
	}

	for _, s := range rawCSVData {
		fmt.Println(s)
	}
}
