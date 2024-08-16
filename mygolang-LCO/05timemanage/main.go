package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(" helloe , welcome to time  handling in golang")
	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006 Monday "))

	// for create a date
	createdDate := time.Date(2020, time.September, 10, 23, 23, 0, 0, time.Local)
	fmt.Println(createdDate)

}
