package main

import "fmt"

// Go has four types of variables -- int , float , string, bool
//with var keyword

func main() {
	var name string = "Mrinal"
	var age int = 24
	code := "Fun"

	fmt.Println("Hello", name, "you are", age, "years old")
	fmt.Printf("Your code is %v", code) //go formatting rules

}
