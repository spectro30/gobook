package main

import "fmt"

func fib(x int) int {
	if x<=1{
		return x
	}
	return fib(x-1) + fib(x-2)

}


func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(fib(n))
}