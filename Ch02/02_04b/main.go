package main

import "fmt"

func main() {
	i1, i2, i3 := 1, 2, 3
	intSum := i1 + i2 + i3
	fmt.Println("Sum of integers:", intSum)
	//fmt.Println("Math")


	f1, f2, f3 := 1.1, 2.2, 3.3
	floatSum := f1 + f2 + f3
	fmt.Println("Sum of floats:", floatSum)

	totalSum := float64(i1) + f2
	fmt.Println("Sum of int and float:", totalSum)

	product := float64(i1) * f2
	fmt.Println("Product:", product)

}
