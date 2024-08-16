package main

import "fmt"

func main() {
	fmt.Println("welcom in pointer")
	// var Ptr *int

	// fmt.Println("value of pointer is ", Ptr)

	Number := 23
	var ptr = &Number

	fmt.Println("value of actual pointer is ", ptr)
	fmt.Println("value of actual pointer is ", *ptr)

	*ptr = *ptr * 2
	fmt.Println("new value is ", Number)
}
