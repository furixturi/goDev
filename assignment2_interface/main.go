package main

import (
	"fmt"
)

type triangle struct {
	b float64
	h float64
}

type square struct {
	a float64
}

type shape interface {
	getArea() float64
}

func main() {
	t := triangle{
		b: 5,
		h: 2,
	}
	printArea(t)

	s := square{
		a: 3,
	}
	printArea(s)
}

func (t triangle) getArea() float64 {
	return t.b * t.h / 2
}

func (sq square) getArea() float64 {
	return sq.a * sq.a
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}
