package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2})

	v := Vertex{3, 5}
	v.X = 4
	fmt.Println(v.X, v.Y)
}
