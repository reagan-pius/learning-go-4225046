package main

import (
	"fmt"
	"time"
)

func main() {

	//fmt.Println("Dates and times")
	t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Printf("Go was launched at %s\n", t)

	n := time.Now()
	fmt.Printf("Current time: %s\n", n)
	fmt.Printf("The object's type: %T\n", n)
	fmt.Printf("The object's value: %v\n", n)


	fmt.Printf(n.Format((time.ANSIC)) + "\n")

	tomorrow := n.AddDate(0, 0, 1)
	fmt.Printf(tomorrow.Format(time.ANSIC) + "\n")
}
