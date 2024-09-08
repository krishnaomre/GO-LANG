package main

import "fmt"

func main() {
	fmt.Println("welcome to aaray in go lang ")
	var fruitlist [5]string
	// compelsary to define a length f arrays
	fruitlist[0] = "appel"
	fruitlist[1] = "mango"
	fruitlist[2] = "banana"
	fmt.Println("fruit list is:", fruitlist)
	fmt.Println("fruit list is:", len(fruitlist))

	var veglist = [8]string{"banan", "paneer", "potato"}
	// veglist[0] = "paneer"
	// veglist[1] = " potato"

	fmt.Println("veg list is:", len(veglist))
}
