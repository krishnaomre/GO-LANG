package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	welcome := "welcome here gyus "

	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("rating  for this :")

	// comma || err err

	input, _ := reader.ReadString('\n')
	fmt.Println("tanks for rating ", input)
	fmt.Printf("type os this rating is  %T", input)

}
