package main

import (
	"fmt"
)

func main() {
	poodle := Dog{"Poodle", 50}
	fmt.Println(poodle)
	fmt.Printf("%+v\n", poodle)
	fmt.Printf("Breed: %v, Weight: %v\n", poodle.Breed, poodle.Weight)
	poodle.Weight = 60
	fmt.Println(poodle)
}

type Dog struct {
	Breed string
	Weight int
}