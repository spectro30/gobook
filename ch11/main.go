package main

import "fmt"
import "gobook/ch11/math"

func main() {
	xs := []float64{1, 2, 3, 4}
	fmt.Println(math.Average(xs))
	fmt.Println(math.Min(xs))
	fmt.Println(math.Max(xs))

}