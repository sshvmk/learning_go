package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11}

	s1 := s[1:4]
	fmt.Println(s1)

	s2 := s[:2]
	fmt.Println(s2)

	s3 := s[1:]
	fmt.Println(s3)

	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)
}
