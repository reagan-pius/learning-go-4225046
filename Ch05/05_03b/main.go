package main

import (
	"fmt"
	"time"
)

func main() {
	go say("Hello from goroutine")
	fmt.Println("Hello from main")

	go func(message string){
		fmt.Println(message)
	}("Hello from anonymous function")
	time.Sleep(2 * time.Second)	
	fmt.Println("All done")
}

func say(message string) {
	time.Sleep(1 * time.Second)
	// fmt.Println(message)
	fmt.Println(message)
}
