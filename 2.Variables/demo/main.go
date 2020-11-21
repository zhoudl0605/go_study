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
    // variable in Go lan must be used, or the compiler will give error 
    fmt.Printf("%s is %d years old, if he is hungry: %t\n", name, age, isHungry)
    fmt.Print("Age: ", age)
    fmt.Println("Name: ", name)
}