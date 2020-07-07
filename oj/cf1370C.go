package main

import "fmt"

func ash(){
	fmt.Println("Ashishgup")
}

func fast(){
	fmt.Println("FastestFinger")
}

func isPrime (x int64) bool{
	if x == 1{
		return false
	}
	if x == 2 {
		return true
	}
	var i int64
	for i=2;(i*i)<=x;i++{
		if x%i == 0 {
			return false
		}
	}
	return true
}

func solve(n int64){
	if n==1 {
		fast()
		return
	}
	if n==2 {
		ash()
		return
	}
	if (n%2==0) && isPrime(n/2) {
		fast()
		return
	}
	for n%2==0{
		n /= 2
	}
	if n == 1{
		fast()
	}else {
		ash()
	}
	return
}


func main() {
	var tc int64
	fmt.Scan(&tc)
	for tc>0 {
		var n int64
		fmt.Scan(&n)
		solve(n)
		tc--
	}
}