package main

import "fmt"

func main() {
	arr := [5]int64{1, 2, 4, 5, 6}
	x := arr[4:4]
	for _, i := range x{
		fmt.Println(i)
	}


}
