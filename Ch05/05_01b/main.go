package main

import (
	"fmt"
)

func main() {
	fmt.Println("Functions")
	doSomething()

	add(4, 5, 6)


	addAllValues(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
}
func doSomething() {
	fmt.Println("Doing something")
} 

func add(value1 int, value2 int, value3 int) int {
	sum := value1 + value2 + value3
	fmt.Printf("Sum: %v\n", sum)
	return sum
}

func addAllValues(values ...int) int{
	sum := 0
	for _, value := range values {
		sum += value
	}
	fmt.Printf("Sum: %v\n", sum)
	return sum
}