package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return // naked return
}

// var c, python, java bool // variable at package level

var i, j int = 1, 2 // variable initialiased at decleration

func main() {
	fmt.Print("Enter two number: ")
	var num1, num2 int
	fmt.Scanln(&num1, &num2)
	fmt.Println(add(num1, num2))

	a, b := swap("first", "second")
	fmt.Println(a, b)

	fmt.Println(split(17))

	// var i int // variable at function level
	// fmt.Println(i, c, python, java)

	// var c, python, java = true, false, "no!" // better way below
	c, python, java := true, false, "no!"
	fmt.Println(i, j, c, python, java)
}
