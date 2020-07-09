package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	res := make([][]uint8, dx)
	for i := range res{
		res[i] = make([]uint8, dy)
	}
	for x:=int(0);x<200;x++{
		for y:=int(0);y<200;y++{
			res[x][y] = uint8((x+y)/2)
		}
	}

	return res
}

func main() {
	pic.Show(Pic)
}
