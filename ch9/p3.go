package main

import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
	peri() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64{
	return r.width * r.height
}

func (c circle) area() float64{
	return math.Pi * c.radius * c.radius
}

func (r rect) peri() float64{
	return 2 * (r.width + r.height)
}

func (c circle) peri() float64{
	return math.Pi * c.radius * 2
}

func solve(s Shape){
	fmt.Println(s)
	fmt.Println(s.area())
	fmt.Println(s.peri())
}

func main() {
	r := rect{3, 4}
	c := circle{5}
	solve(r)
	solve(c)
}