//	3.Write a program that defines an interface type called Shape with two methods:
//	Area() and Perimeter(), both of which return a float64 value.
//	Implement this interface for three different types: Rectangle, Circle, and Triangle.

package main

import (
	"fmt"
	"math"
)

type shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	lenght, width float64
}

type Circle struct {
	radius float64
}

type Triangle struct {
	height, base float64
}

func (r Rectangle) Area() float64 {
	return r.lenght * r.width
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (t Triangle) Area() float64 {
	return 0.5 * t.height * t.base
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.lenght + r.width)
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (t Triangle) Perimeter() float64 {
	// consoidering its a isosceles triangle
	side := math.Sqrt((t.base/2)*(t.base/2) + t.height*t.height)
	return 2*side + t.base
}

func getInfo(s shape) {
	fmt.Printf("Shape : %T,\t Area : %v,\t Perimeter : %v \n", s, s.Area(), s.Perimeter())
}

func main() {

	r1 := Rectangle{12, 30}
	c1 := Circle{20}
	t1 := Triangle{4, 6}

	getInfo(r1)
	getInfo(c1)
	getInfo(t1)
}
