package main

import "fmt"

func main() {
	const PI = 3.141592653589793
	fmt.Println(PI)

	// const has 2 types in Go -- Typed const & Untyped const
	// Typed const
	const Mk int = 1
	fmt.Println(Mk)

	// Untyped const
	const Mk2 = 7
	fmt.Println(Mk2)

	//we also can use multiple consts in one line

}
