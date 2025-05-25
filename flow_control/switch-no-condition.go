package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Printf("%s\n", t)

	switch {
	case t.Hour() < 12:
		fmt.Println("it is morning")
	case t.Hour() < 17:
		fmt.Println("it is afternoon")
	default:
		fmt.Println("it is evening")
	}
}
