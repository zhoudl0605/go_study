package main

import (
	"fmt"
	"strings"
)

func main() {
	// \ has special meaning
	path0 := "E:\\Riot Games\\League of Legends"
	fmt.Println(path0)

	// use double quotes
	path1 := "\"E:\\Riot Games\\League of Legends\""
	fmt.Println(path1)

	path2 := "'E:\\Riot Games\\League of Legends'"
	fmt.Println(path2)

	// check length
	fmt.Println(len(path1))

	// splice string together
	name := "Alex"
	gender := "male"

	s1 := name + " " + gender
	fmt.Println(s1)

	s2 := fmt.Sprintf("%s %s", name, gender)
	fmt.Println(s2)

	// Seperate
	ret := strings.Split(path0, "\\")
	fmt.Println(ret)

	// Contains
	fmt.Println(strings.Contains(s1, "Alex"))
	fmt.Println(strings.Contains(s1, "Mike"))

	// Prefix
	fmt.Println(strings.HasPrefix(s1, "Alex"))
	fmt.Println(strings.HasPrefix(s1, "Mike"))

	//Suffix
	fmt.Println(strings.HasSuffix(s1, "Alex"))
	fmt.Println(strings.HasSuffix(s1, "male"))

	// Join
	fmt.Println(strings.Join(ret, "\\"))
}
