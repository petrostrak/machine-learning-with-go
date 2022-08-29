// A vector is an ordered collection of numbers arranged in either a row or
// column. Each of the numbers in a vector is called a component. This might
// be, for example, a collection of numbers that represents our company sales,
// or it might be a collection of numbers representing temperatures.
package main

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

func main() {

	// Represent a vector in Go as a slice
	var vec []float64

	// Add components to the vector.
	vec = append(vec, 11.0)
	vec = append(vec, 5.2)

	// Output the result to stdout.
	for _, v := range vec {
		fmt.Println(v)
	}

	// Create a new vector
	vector := mat.NewVecDense(2, vec)

	fmt.Println(vector)
}
