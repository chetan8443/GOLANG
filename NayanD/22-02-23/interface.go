package main

import "fmt"

type shape interface {
	area() float64
}
type circle struct {
	radius float64
}
type square struct {
	side float64
}

func (c circle) area() float64 {
	return 3.14 * c.radius * c.radius

}

func (s square) area() float64 {
	return s.side * s.side

}

func getArea(s shape) {
	fmt.Printf("Area : %v \n", s.area())
}

func main() {

	c := circle{10}
	getArea(c)
	sq := square{5}
	getArea(sq)

}
