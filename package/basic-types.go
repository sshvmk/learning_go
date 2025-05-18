package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

// variable declarations may be "factored" into blocks
// as with import statements.
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	// zero value
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s)

	// type conversion
	var x, y int = 3, 4
	var a float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(a)
	fmt.Println(x, a, z)

	v := true // change RHS and see changes when printing
	fmt.Printf("v is of type %T\n", v)
}
