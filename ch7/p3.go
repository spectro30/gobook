package main

import "fmt"

func max(args ...int) int {
	mx  := 0
	for _, i := range args{
		if i >= mx{
			mx = i
		}
	}
	return mx
}

func main() {
	fmt.Println(max(10, 100, 1, 5, 101, 5))
}
