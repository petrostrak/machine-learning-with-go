package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

// CSVRecord contains a successfully parsed row
// of the CSV file.
type CSVRecord struct {
	SepalLength float64
	SepalWidth  float64
	PetalLength float64
	PetalWidth  float64
	Species     string
	ParseError  error
}

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

	// Create a slice of CSVRecord
	var csvData []CSVRecord

	for {
		// Read in a row
		record, err := r.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Create a CSVRecord var for each row.
		var csvRecord CSVRecord

		// Parse each of the values in the record
		// based on an expected type.
		for i, value := range record {
			// Parse the value in the record as a string
			// for the string column.
			if i == 4 {
				// Check whether is empty or not.
				if value == "" {
					fmt.Printf("Unexpected type in column %d\n", i)
					csvRecord.ParseError = fmt.Errorf("empty string value")
					break
				}

				// Add the string value to the VSCRecord.
				csvRecord.Species = value
				continue
			}

			// Otherwise, parse the value in the record as a float64.
			var floatValue float64

			// If the value cannot be parsed
			if floatValue, err = strconv.ParseFloat(value, 64); err != nil {
				fmt.Println(err)
				csvRecord.ParseError = fmt.Errorf("could not parse float")
				break
			}

			// Add the float value to the respective field in the
			// CSVRecord.
			switch i {
			case 0:
				csvRecord.SepalLength = floatValue
			case 1:
				csvRecord.SepalWidth = floatValue
			case 2:
				csvRecord.SepalLength = floatValue
			case 3:
				csvRecord.SepalWidth = floatValue
			}
		}

		// Append successfully parsed records to the slice defined.
		if csvRecord.ParseError == nil {
			csvData = append(csvData, csvRecord)
		}
	}

	for _, s := range csvData {
		fmt.Println(s)
	}
}
