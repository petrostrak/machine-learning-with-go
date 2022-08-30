package main

import (
	"fmt"
	"log"
	"math"

	"gonum.org/v1/gonum/mat"
)

func main() {

	// Create two matrices of the same size
	a := mat.NewDense(3, 3, []float64{1, 2, 3, 0, 4, 5, 0, 0, 6})
	b := mat.NewDense(3, 3, []float64{8, 9, 10, 1, 4, 2, 9, 0, 2})

	// Create a third matrix of different size.
	c := mat.NewDense(3, 3, []float64{3, 2, 1, 0, 8, 5, 0, 6, 1})

	// Add a and b
	d := mat.NewDense(3, 3, nil)
	d.Add(a, b)
	formatted := mat.Formatted(d)
	fmt.Println(formatted)
	fmt.Println()

	// Multiply a and c.
	e := mat.NewDense(3, 3, nil)
	e.Mul(a, c)
	formatted = mat.Formatted(e)
	fmt.Println(formatted)
	fmt.Println()

	// Raising a matrix to a power.
	f := mat.NewDense(3, 3, nil)
	f.Pow(a, 3)
	formatted = mat.Formatted(f)
	fmt.Println(formatted)
	fmt.Println()

	// Apply a function to each of the elements of a.
	g := mat.NewDense(3, 3, nil)
	sqrt := func(_, _ int, v float64) float64 { return math.Sqrt(v) }
	g.Apply(sqrt, a)
	formatted = mat.Formatted(g)
	fmt.Println(formatted)
	fmt.Println()

	// Compute and output the transpose of the matrix.
	formatted = mat.Formatted(a.T())
	fmt.Println(formatted)
	fmt.Println()

	// Compute and output the determinant of a.
	deta := mat.Det(a)
	fmt.Printf("det(a) = %.2f\n\n", deta)
	fmt.Println()

	// Compute and output the inverse of a.
	aInverse := mat.NewDense(3, 3, nil)
	if err := aInverse.Inverse(a); err != nil {
		log.Fatal(err)
	}
	formatted = mat.Formatted(aInverse)
	fmt.Printf("a^-1 = \n%v\n\n", formatted)
}
