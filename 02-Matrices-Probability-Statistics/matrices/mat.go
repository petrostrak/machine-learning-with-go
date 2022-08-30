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
}
