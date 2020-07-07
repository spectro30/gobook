package main

import "fmt"

func main() {
	arr := [5]float64{1, 2, 10, 4, 2}
	var sum float64 = 0

	for _, x := range arr {
		sum += x
	}
	fmt.Println(sum / 5)


}
