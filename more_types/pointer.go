package main

import "fmt"

func main() {
	i, j := 42, 690

	p := &i         // p is a pointer to i
	fmt.Println(*p) // read i through pointer
	*p = 21         // change value of i using pointer
	fmt.Println(i)  // read i to see new value reflect

	q := &j
	*q = *q / 69
	fmt.Println(j)
}
