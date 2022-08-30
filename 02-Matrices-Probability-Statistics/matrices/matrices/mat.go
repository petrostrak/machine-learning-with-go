package main

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

// To form a matrix, we need to create a slice of float64 values that
// is a flat representation of all the matrix components.
func main() {

	// Create a flat representation of our matrix.
	data := []float64{1.2, -5.7, -2.4, 7.3}

	// Form our matrix. The first argument is the number of
	// rows, and the second the number of columns.
	a := mat.NewDense(2, 2, data)

	formatted := mat.Formatted(a)
	fmt.Println(formatted)

	// Get a single value from the matrix.
	val := a.At(0, 1)
	fmt.Printf("The value of a at (0,1) is: %.2f\n\n", val)

	// Get the values on a specific column.
	col := mat.Col(nil, 0, a)
	fmt.Printf("The values in the 1st column are: %v\n\n", col)

	// Get the values in a specific row.
	row := mat.Row(nil, 1, a)
	fmt.Printf("The values in the 2nd row are: %v\n\n", row)

	// Modify a single element.
	a.Set(0, 1, 11.2)

	// Modify an entire row.
	a.SetRow(0, []float64{14.3, -4.1})

	// Modify an entire column.
	a.SetCol(0, []float64{1.7, -0.3})
	fmt.Println(formatted)
}
