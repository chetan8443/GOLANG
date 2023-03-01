// Shape interface with implementation for various Shape types
package main

import (
	"fmt"
	"math"
)

// Shape is the interface for Area() and Perimeter() functions
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Rectangle Struct can have a length and width variable. Accordingly, we'll be implementing functions
type Rectangle struct {
	length float64
	width  float64
}

func (r Rectangle) Area() float64 {
	// l * w
	return r.length * r.width
}

func (r Rectangle) Perimeter() float64 {
	// 2(l + w)
	return 2 * (r.length + r.width)
}

func runRectangleUtils() {
	var myRectangle Rectangle
	myRectangle = Rectangle{3, 4}
	fmt.Println("Rectangle area is: ", myRectangle.Area())
	fmt.Println("Rectangle Perimeter is: ", myRectangle.Perimeter())
}

// Circle Struct can only have a radius variable. Accordingly, we'll be implementing functions
type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	// pie * r^2
	return 3.14 * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	// 2 * pie * r
	return 2 * 3.14 * c.radius
}

func runCircleUtils() {
	var myCircle Circle
	myCircle = Circle{5}
	fmt.Println("Circle area is: ", myCircle.Area())
	fmt.Println("Circle Perimeter is: ", myCircle.Perimeter())
}

// Triangle Struct can only have a radius variable. Accordingly, we'll be implementing functions
type Triangle struct {
	height float64
	base   float64
}

// Implementing Triange functions
func (t Triangle) Area() float64 {
	// (h * b) / 2
	return (t.height * t.base) / 2
}

func (t Triangle) Perimeter() float64 {
	hypotenuse := math.Sqrt(t.height*t.height + t.base*t.base)
	return (hypotenuse + t.height + t.base)
}

func runTriangleUtils() {
	var myTriangle Triangle
	myTriangle = Triangle{3, 4}
	fmt.Println("Triangle area is: ", myTriangle.Area())
	fmt.Println("Triangle Perimeter is: ", myTriangle.Perimeter())
}

func main() {
	runRectangleUtils()
	runCircleUtils()
	runTriangleUtils()
}
