package main

import "fmt"

func solve (){
	var n int64
	fmt.Scan(&n)
	fmt.Println(n/2)
}

func main() {
	var tc int64
	fmt.Scan(&tc)
	for tc>0 {
		solve()
		tc--
	}
}