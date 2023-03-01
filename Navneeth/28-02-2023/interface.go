//	3.Write a program that defines an interface type called Shape with two methods: Area() and Perimeter(), both of which return a float64 value. Implement this interface for three different types: Rectangle, Circle, and Triangle.

package main

import (
	"fmt"
	"math"
)

func main() {

	var l, w, r, h, b float64

	fmt.Println("Enter length and width for rectangle: ")
	fmt.Scanf("%f %f", &l, &w)
	fmt.Println("Enter radius for circle: ")
	fmt.Scanf("%f", &r)
	fmt.Println("Enter height , base for triangle: ")
	fmt.Scanf("%f %f", &h, &b)

	r1 := Rectangle{l, w}
	c1 := Circle{r}
	t1 := Triangle{h, b}

	getInfo(r1)
	getInfo(c1)
	getInfo(t1)
}

type shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	l, w float64
}

type Circle struct {
	r float64
}

type Triangle struct {
	h, b float64
}

func (r Rectangle) Area() float64 {
	return r.l * r.w
}

func (c Circle) Area() float64 {
	return math.Pi * c.r * c.r
}

func (t Triangle) Area() float64 {
	return 0.5 * t.h * t.b
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.l + r.w)
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.r
}

func (t Triangle) Perimeter() float64 {
	side := math.Sqrt((t.b/2)*(t.b/2) + t.h*t.h)
	return 2*side + t.b
}

func getInfo(s shape) {
	fmt.Printf("Shape : %T,\t Area : %v,\t Perimeter : %v \n", s, s.Area(), s.Perimeter())
}

