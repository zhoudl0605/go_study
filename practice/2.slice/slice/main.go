package main

import "fmt"

func main() {
	// slice
	var s1 []int    // determin a slice to save int, with unlimit size
	var s2 []string // determin a slice to save string, with unlimit size
	fmt.Println(s1, s2)

	// initialize
	s1 = []int{1, 2, 3}
	s2 = []string{"hello", "world"}
	fmt.Println(s1, s2)
	fmt.Println(s1 == nil) // false, check if s1 is empty
	fmt.Println(s2 == nil) // false, check if s2 is empty
	fmt.Printf("s1 length: %d, capacity: %d\n", len(s1), cap(s1))
	fmt.Printf("s2 length: %d, capacity: %d\n", len(s2), cap(s2))

	// get slice from array
	a1 := [...]int{1, 3, 5, 7, 9, 11, 13}
	s3 := a1[0:4] // [1 3 5 7] get slice from array
	fmt.Println(s3)
	s4 := a1[1:6] // [3 5 7 9 11]
	fmt.Println(s4)
	s5 := a1[:4] // [1 3 5 7]
	fmt.Println(s5)
	s6 := a1[3:] // [7 9 11 13]
	fmt.Println(s6)
	s7 := a1[:] //	[1 3 5 7 9 11 13]
	fmt.Println(s7)

	// capacity of slice is the capacitiy of the base array from the beginning index to the end index
	// capacity is starting from 1 to 13, equal to 7
	fmt.Printf("len(s5): %d, cap(s5): %d\n", len(s5), cap(s5))
	// capacity is starting from 7 to 13, equal to 4
	fmt.Printf("len(s6): %d, cap(s6): %d\n", len(s6), cap(s6))

	// slice of slice
	s8 := s6[3:] // [13]
	fmt.Printf("len(s8): %d, cap(s8): %d\n", len(s8), cap(s8))

	// slice is a reference to object, reference to an array
	fmt.Println("s6: ", s6)
	a1[6] = 1000 // change the value in the base array
	fmt.Println("s6: ", s6)
}
