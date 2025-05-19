package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0

	for i := 0; i < 10; i++ { // Initial: 10 iterations (for demonstration)
		prevZ := z
		z -= (z*z - x) / (2 * z)                     // Newton's method
		fmt.Printf("Iteration %d: z = %f\n", i+1, z) // Print z at each iteration

		if math.Abs(z-prevZ) < 1e-6 { // Check for convergence (change less than 1e-6)
			fmt.Println("Converged after", i+1, "iterations.")
			break // Exit the loop if converged
		}
	}
	return z
}

func main() {
	for i := 1.0; i <= 5.0; i++ {
		fmt.Printf("Sqrt(%f) = %f\n", i, Sqrt(i))
		fmt.Printf("math.Sqrt(%f) = %f\n", i, math.Sqrt(i)) // Compare with standard library
		fmt.Printf("Difference: %e\n\n", Sqrt(i)-math.Sqrt(i))
	}

	x := 25.0
	mySqrt := math.Sqrt(x)
	fmt.Printf("The square root of %f is approximately %f\n", x, mySqrt)

	x = 2
	mySqrt = math.Sqrt(x)
	fmt.Printf("The square root of %f is approximately %f\n", x, mySqrt)
}
