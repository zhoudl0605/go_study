package main

import "fmt"

// constant
const (
	pi = 3.1415926
	e  = 2.7182
)

const (
	n1 = 100
	n2
	n3
)

const (
	m1 = iota // 0
	m2        // 1
	m3        // 2
	m4        // 3
)

const (
	d1, d2 = iota + 1, iota + 2 // d1 = 1, d2 = 2
	d3, d4 = iota + 1, iota + 2 // d3 = 3, d4 = 4
)

const (
	b1 = iota // 0
	b2        //1
	_
	b3 // 3
)

func main() {
	fmt.Println(pi)
	// pi = 1234 // error
	fmt.Println(n1, n2, n3)
	fmt.Println(m1, m2, m3, m4)
	fmt.Println(b1, b2, b3)
	fmt.Println(d1, d2, d3, d4)
}
