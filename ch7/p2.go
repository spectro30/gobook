package main

import "fmt"

func half (x int64) bool{
	x /= 2
	if x % 2 == 1 {
		return true
	}
	return false

}
func main() {
	fmt.Println(half(2))
}
