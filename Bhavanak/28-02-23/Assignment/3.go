package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
	perimeter() float64
}

type rectangle struct {
	side1 float64
	side2 float64
}

type triangle struct {
	base   float64
	height float64
	side   float64
}
type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (r rectangle) area() float64 {
	return r.side1 * r.side2
}

func (r rectangle) perimeter() float64 {
	return 2*r.side1 + r.side2
}

func (t triangle) area() float64 {
	return 0.5 * t.base * t.height
}

func (t triangle) perimeter() float64 {
	return t.base + t.height + t.side
}

func calculate(s shape) {
	fmt.Println("Area:", s.area())
	fmt.Println("Perimeter:", s.perimeter())
}

func main() {
	c1 := circle{8}
	r1 := rectangle{14, 12}
	t1 := triangle{2, 4, 6}

	fmt.Println("Circle")
	calculate(c1)
	fmt.Println("Rectangle")
	calculate(r1)
	fmt.Println("Triangle")
	calculate(t1)

}
