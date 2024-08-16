package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("hello gyus, welcome pizza app")
	fmt.Println("please rate our pizza between 1 to 5")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')
	fmt.Println("tanks for rating ,", input)

	// fmt.Println("thanks for rating ")

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)

	} else {
		fmt.Println("added 1 to your rating:", numRating+1)
	}
}
