package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a [2]string
	a[0] = "Learning"
	a[1] = "Go"
	fmt.Println(a[0], a[1])
	fmt.Println(a)
	fmt.Println(reflect.TypeOf(a))

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
	fmt.Println(reflect.TypeOf(primes))

	// slicing in array
	var s []int = primes[1:4]
	fmt.Println(s)
}
