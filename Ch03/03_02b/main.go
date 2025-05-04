package main

import (
	"fmt"
)

func main() {
	anInt := 42
	var p *int = &anInt // p points to anInt
	if p == nil {
		fmt.Println("p is nil")
	} else {
		fmt.Println("The value of p:", *p)
	}

	value1 := 42.13
	pointer1 := &value1 // ampersand operator (&) gives the address of value1
	*pointer1 = *pointer1/31
	fmt.Println(("Pointer: "), *pointer1)
	fmt.Println("The original value:", value1)
}
