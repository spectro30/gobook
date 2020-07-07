package main

import "fmt"

func main() {

	fmt.Println("Enter length in feet: ")
	var feet float64
	fmt.Scanf("%f", &feet)
	meter := 0.3048 * feet
	fmt.Println(meter)
}
