package main

import (
	"fmt"
)

func main() {
	f1 := 1.23456
	fmt.Printf("%T", f1) // by default the decimal is float64

	f2 := float32(1.23456)
	fmt.Printf("%T", f2)
	// f1 = f2 // float 32 can not assign to float64
}
