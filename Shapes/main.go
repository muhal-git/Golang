package main

import "fmt"

type square struct {
	sideLength float64
}
type triangle struct {
	base   float64
	heigth float64
}

type shape interface {
	getArea() float64
}

func main() {

	t := triangle{base: 11, heigth: 5}
	s := square{sideLength: 5}

	printArea(t)
	printArea(s)
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}
func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.heigth
}
func printArea(s shape) {
	fmt.Println(s.getArea())
}
