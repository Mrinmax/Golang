package main

import "fmt"

func main() {
	//Go has three different function to output text--Print(), Println(), Printf()
	//Print() is used to output a single line of text
	//print() does not add a newline at the end of the output & space is not added between the output
	var x, y string = "hello", "world"
	fmt.Print(x, " ", y)

	//Println() is used to output a single line of text & add a newline at the end of the output
	// fmt.Println(x, y)

	//printf() is used to formatting the output
	// %v & %T are used to display the value and type of the variables

	fmt.Printf("%v %T\n", x, x)
	fmt.Printf("%v %T\n", y, y)
}
