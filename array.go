package main

import "fmt"

func main() {
	var arr1 = [3]int{1, 2, 3}
	arr2 := [5]int{4, 5, 6, 7, 8}

	arr2[3] = 100

	fmt.Println(arr1[2])
	fmt.Println(arr2)

	//finding the length of an array
	fmt.Println(len(arr2))
}
