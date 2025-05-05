package main

import (
	"fmt"
	"time"
)

func main() {

	weekday := time.Now().Weekday()
	fmt.Printf("Today is %v\n", weekday)
	dayNumber := 6
	// dayNumber := int(weekday)
	fmt.Printf("The day as a number is %v\n", dayNumber)

	var result string
	switch dayNumber {
	case 1:
		result = "Its a Monday!"
	case 2:
		result = "Its a Tuesday!"
	case 3:
		result = "Its a Wednesday!"
	case 4:
		result = "Its a Thursday!"
	case 5:
		result = "Its a Friday!"
	default:
		result = "Its a weekend!"	
}
fmt.Println(result)
}
