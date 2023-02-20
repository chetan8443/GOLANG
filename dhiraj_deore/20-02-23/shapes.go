// Write a funcction to to print area of different shapes using interface

package main

import "fmt"

type shape interface {
	area() float64
}

type circle struct {
	redius float64
}

type square struct {
	side float64
}

func (c circle) area() float64 {
	return 3.14 * c.redius * c.redius
}

func (s square) area() float64 {
	return s.side * s.side
}

func getArea(s shape) {
	fmt.Printf("Area : %v \n", s.area())
}

func main() {
	c := circle{20}
	getArea(c)

	sq := square{20}
	getArea(sq)
}
