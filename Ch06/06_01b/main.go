package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fileName := "./fromString.txt"
	file, err := os.Create(fileName)
	defer file.Close()
	checkErr(err)
	length, err := io.WriteString(file, "Hello, World!")
	fmt.Printf("Wrote a file with %v characters\n", length)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func readFile(fileName string) {
	file, err := os.ReadFile(fileName)
	checkErr(err)
	fmt.Println("Text read from file:", string(file))

}