package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	mp := make(map[string]int)
	words := strings.Fields(s)
	for _, x:= range words{
		mp[x] ++
	}
	return mp
}

func main() {
	wc.Test(WordCount)
}
