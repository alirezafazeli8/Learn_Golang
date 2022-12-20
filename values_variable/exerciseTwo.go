package main

import "fmt"

func main() {
	var numberPI float32 = 3.14
	var circleRadius int = 5

	sum := 2 * numberPI * float32(circleRadius)

	fmt.Println(sum)
	fmt.Printf("For a radius of %v, the circle circumference is %v", circleRadius, sum)

}
