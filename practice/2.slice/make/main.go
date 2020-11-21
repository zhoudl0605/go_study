package main

import "fmt"

// use make() to create a slice

func main() {
	s1 := make([]int, 5, 10)
	fmt.Printf("s1 = %v len(s1)=%d cap(s1)=%d \n", s1, len(s1), cap(s1))

	s2 := make([]int, 0, 10)
	fmt.Printf("s2 = %v len(s2)=%d cap(s2)=%d \n", s2, len(s2), cap(s2))
	fmt.Printf("%t", true)
}
