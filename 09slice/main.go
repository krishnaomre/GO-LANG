package main

import (
	"fmt"
)

func main() {
	fmt.Println("welcome to slice in golang ")
	var fruitlist = []string{"apple", "potato", "tomato"}
	fmt.Println("type of fruitlis is:", fruitlist)
	fmt.Println("here is fruitlist:", fruitlist)
	fruitlist = append(fruitlist, "peach")
	fmt.Println("new fruitlist", fruitlist)

	var songlist = []string{"tuba-tuba", "mombattiya", "kinni sonni"}
	fmt.Println("songlist", songlist)
	songlist = append(songlist, "fasla", "mahiye jinna sona", " mehrma ")
	fmt.Println(" new songlist", songlist)
	fmt.Println("length", len(songlist))
	songlist = append(songlist[1:])
	fmt.Println(" fruitlist", songlist)

}
