package main

import "fmt"

var(
	name string
    age int
    isHungry bool
)

func main(){
    name = "Alex Zhou"
    age = 23
    isHungry = true
    
    fmt.Printf("%s is %d years old, if he is hungry: %t\n", name, age, isHungry)
}