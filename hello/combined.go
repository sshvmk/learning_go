package main

import "fmt"

func main() {
	fmt.Print("Enter a number: ")
	var num int
	fmt.Scanln(&num) // Reads input
	fmt.Printf("You entered: %d\n", num)
	fmt.Println("Double:", num*2)
}
