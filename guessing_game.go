package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())
	target := rand.Intn(100) + 1 // Random number between 1 and 100

	fmt.Println("🎮 Welcome to the Number Guessing Game!")
	fmt.Println("🔢 I've picked a number between 1 and 100. Can you guess it?")

	var guess, attempts int

	for {
		fmt.Print("👉 Enter your guess: ")
		_, err := fmt.Scanln(&guess)
		if err != nil {
			fmt.Println("⚠️ Invalid input! Please enter a number.")
			// Clear input buffer to avoid infinite loop
			var discard string
			fmt.Scanln(&discard)
			continue
		}

		attempts++

		if guess < target {
			fmt.Println("⬆️ Too low! Try a higher number.")
		} else if guess > target {
			fmt.Println("⬇️ Too high! Try a lower number.")
		} else {
			fmt.Printf("🎉 Correct! You guessed the number in %d attempts!\n", attempts)
			break
		}
	}
}
