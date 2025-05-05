package main

import (
	"fmt"
)

func main() {
	// theAnswer := -1
	var result string

	if theAnswer := 42; theAnswer < 0 {
		result = "less than zero"
	} else if theAnswer == 0 {
		result = "equal to zero"
	} else {
		result = "greater than zero"
	}
	fmt.Println(result)
}
