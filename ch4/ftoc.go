package main

import "fmt"

func main() {

	fmt.Println("Enter Fahrenheit temperature: ")
	var far float64
	fmt.Scanf("%f", &far)
	cel := ( (far - 32) * 5) / 9
	fmt.Println(cel)
}
