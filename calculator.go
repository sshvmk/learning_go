package main

import (
	"fmt"
	"os"
)

// Add returns the sum of two numbers
func Add(a, b float64) float64 {
	return a + b
}

// Subtract returns the difference of two numbers
func Subtract(a, b float64) float64 {
	return a - b
}

// Multiply returns the product of two numbers
func Multiply(a, b float64) float64 {
	return a * b
}

// Divide returns the quotient of two numbers (checks for division by zero)
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

func main() {
	var num1, num2 float64
	var operator string

	fmt.Print("Enter first number: ")
	_, err := fmt.Scanln(&num1)
	if err != nil {
		fmt.Println("Invalid input!")
		return
	}

	fmt.Print("Enter operator (+, -, *, /): ")
	_, err = fmt.Scanln(&operator)
	if err != nil {
		fmt.Println("Invalid operator!")
		return
	}

	fmt.Print("Enter second number: ")
	_, err = fmt.Scanln(&num2)
	if err != nil {
		fmt.Println("Invalid input!")
		return
	}

	var result float64
	var calcError error

	switch operator {
	case "+":
		result = Add(num1, num2)
	case "-":
		result = Subtract(num1, num2)
	case "*":
		result = Multiply(num1, num2)
	case "/":
		result, calcError = Divide(num1, num2)
		if calcError != nil {
			fmt.Println("Error:", calcError)
			os.Exit(1)
		}
	default:
		fmt.Println("Invalid operator!")
		os.Exit(1)
	}

	fmt.Printf("Result: %.2f %s %.2f = %.2f\n", num1, operator, num2, result)
}
