package main

import "fmt"

/*  array practice
 *	1. find add all elements in the array [1,3,5,7,8]
 *	find index of elements when they add together get 8
 *	such as (0,3) 1 + 7 = 8, (1,2) 3 + 5 = 8
 */

func main() {
	array := []int{1, 3, 5, 7, 8}
	result := 0

	// find the result
	// reuslt = 24
	for _, v := range array {
		result = result + v
	}
	fmt.Println(result)

	// find the elements add can get 8
	for i, v := range array {
		for j := i + 1; j < len(array); j++ {
			if v+array[j] == 8 {
				fmt.Printf("(%d, %d)\n", v, j)
			}
		}
	}
}
