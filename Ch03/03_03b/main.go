package main

import (
	"fmt"
)

func main() {
	var colors [3]string
	colors[0] = "red"
	colors[1] = "green"
	colors[2] = "blue"
	fmt.Println(colors)

	var numbers = [5]int{5, 3, 2, 4, 1}
	fmt.Println(numbers)
	fmt.Println("The length of colors:", len(colors))
	fmt.Println("The length of numbers:", len(numbers))
}
