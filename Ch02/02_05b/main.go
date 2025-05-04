package main

import (
	"fmt"
	"math"
)

func main() {

	f1, f2, f3 := 23.5, 65.1, 76.3
	sum := f1 + f2 + f3
	fmt.Println("Float sum:", sum)

	sum = math.Round(sum*100) / 100
	fmt.Println("The sum is now \n", sum)

	fmt.Println(("the value of pi is"), math.Pi)

	circleradius := 5.0
	circumference := 2 * math.Pi * circleradius
	fmt.Println("Circumference of circle with radius", circleradius, "is %.2f\n", circumference)

}
