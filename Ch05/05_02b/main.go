package main

import (
	"fmt"
)

func main() {
	dog := Dog{"Poodle", "Woof"}
	dog.Bark()
}


type Dog struct {
   Breed string
   Sound string
}

func (d Dog) Bark() {
	 fmt.Printf("The %v barks: %v!\n", d.Breed, d.Sound)
}