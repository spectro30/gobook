package main

import "fmt"

func zero (xPtr *int){
	*xPtr = 5
}

func main() {
	x := 5
	zero(&x)
	fmt.Println(x)
}