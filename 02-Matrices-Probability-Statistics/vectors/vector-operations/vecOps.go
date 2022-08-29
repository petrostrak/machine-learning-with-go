package main

import (
	"fmt"

	"gonum.org/v1/gonum/blas/blas64"
	"gonum.org/v1/gonum/floats"
	"gonum.org/v1/gonum/mat"
)

func main() {

	// Initialize a couple of vectors represented as slices.
	vectorA := []float64{11.0, 5.2, -1.3}
	vectorB := []float64{-7.2, 4.2, 5.1}

	// Compute the dot product of A and B.
	// (https://en.wikipedia.org/wiki/Dot_product).
	dotProduct := floats.Dot(vectorA, vectorB)
	fmt.Printf("The dot product of A and B is: %0.2f\n", dotProduct)

	// Scale each element of A by 1.5
	floats.Scale(1.5, vectorA)
	fmt.Printf("Scaling A by 1.5 gives: %v\n", vectorA)

	// Compute the norm/length of B.
	normB := floats.Norm(vectorB, 2)
	fmt.Printf("The norm/length of B is: %0.2f\n", normB)

	// Initialize a couple of "vectors" represented as slices
	// using the gonum.org/v1/gonum/mat library.
	vectorC := mat.NewVecDense(3, []float64{11.0, 5.2, -1.3})
	vectorD := mat.NewVecDense(3, []float64{-7.2, 4.2, 5.1})
	// Compute the dot product of A and B
	// (https://en.wikipedia.org/wiki/Dot_product).
	dotProduct = mat.Dot(vectorC, vectorD)
	fmt.Printf("The dot product of C and D is: %0.2f\n", dotProduct)
	// Scale each element of A by 1.5.
	vectorC.ScaleVec(1.5, vectorC)
	fmt.Printf("Scaling C by 1.5 gives: %v\n", vectorA)
	// Compute the norm/length of B.
	normB = blas64.Nrm2(vectorD.RawVector())
	fmt.Printf("The norm/length of D is: %0.2f\n", normB)
}
