package main

import (
	"fmt"
	"sort"
)

func main() {
	// This is a slice
	//var colors = [3]string{"Red", "Green", "Blue"}

	var colors = make([]string, 0, 3)
	colors = append(colors, "Red", "Green", "Blue")
	colors = append(colors, "Yellow", "Black", "White")
	fmt.Println(colors)
	fmt.Println("The length of colors:", len(colors))
	// fmt.Println("The capacity of colors:", cap(colors))
	// The length of the slice is the number of elements in it

	colors = remove(colors, 0)
	// fmt.Println(colors)

	sort.Strings(colors)
	fmt.Println(colors)
	colors = remove(colors, 0)
	fmt.Println(colors)
}

func remove(slice []string, i int) []string {
	return append(slice[:i], slice[i+1:]...)
}
