package main

import "fmt"

// this is a comment
func main() {

	fmt.Print("Enter a number ")
	var input float64
	var name = "Alif"
	fmt.Scanf("%f", &input)
	output := input * 2.0
	fmt.Println(output)
	fmt.Println(name)
}
