package main

import "fmt"

// integer

func main() {
	var i1 = 101
	fmt.Printf("%d\n", i1)
	fmt.Printf("%b\n", i1) // decimal display in binary
	fmt.Printf("%o\n", i1) // decimal display in Octal
	fmt.Printf("%x\n", i1) // decimal display in Hexa

	// Octal integer
	i2 := 077
	fmt.Printf("%d\n", i2)

	// Hexa
	i3 := 0x1234567
	fmt.Printf("%d\n", i3)

	// check variable type
	fmt.Printf("%T\n", i3)

	// declare int8
	i4 := int8(9) // declare the type or will be int
	fmt.Printf("%T\n", i4)
}
