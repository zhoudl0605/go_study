package main

import "fmt"

var(
	name string
    age int
    isHungry bool
)

func foo()(int, string){
	return 10, "abcde"
}

func main(){
    name = "Alex Zhou"
    age = 23
    isHungry = true
    // unglobal variable in Go lan must be used, or the compiler will give error 
    fmt.Printf("%s is %d years old, if he is hungry: %t\n", name, age, isHungry)
    fmt.Print("Age: ", age)
    fmt.Println("Name: ", name)

    // Determine and initializing at the same time
    var s1 string = "Alex Zhou" // not suggest to use this
    fmt.Println(s1)

    // Type Derivation
    var s2 = "Alex"
    fmt.Println(s2)

    // Short variable declaration
    // only able to use inside of the function
    s3 := "Acadia University"
    fmt.Println(s3)

    x,_ := foo()
    fmt.Println(x)
    _,y := foo()
    fmt.Println(y)
}