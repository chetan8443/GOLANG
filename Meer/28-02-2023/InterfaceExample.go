// program that defines an interface type called Shape with two methods: Area()
// and Perimeter(), both of which return a float64 value. Implement this interface for three
// different types: Rectangle, Circle, and Triangle.
package main

import (
	"fmt"
	"math"
)

// creating interface
type Shape interface {
	Area() float64
	Perimeter() float64
}

// creating struct of Rectangle
type Rectangle struct {
	Length, Width float64
}

// creating Function to calculate Area of Rectangle
func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}

// creating Function to calculate Parameter of Rectangle
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Length + r.Width)
}

// creating struct of Circle
type Circle struct {
	Radius float64
}

// creating Function to calculate Area of Circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// creating Function to calculate Parameter of Circle
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// creating struct of Triangle
type Triangle struct {
	height, base float64
}

// creating Function to calculate Area of Triangle
func (t Triangle) Area() float64 {
	return 0.5 * t.height * t.base
}

// creating Function to calculate Parameter of Triangle
func (t Triangle) Perimeter() float64 {
	// consoidering its a isosceles triangle
	side := math.Sqrt((t.base/2)*(t.base/2) + t.height*t.height)
	return 2*side + t.base
}

func main() {
	var s Shape = Circle{5.0}

	fmt.Printf("Circle \t\t: Area = %f, Perimeter = %f\n", s.Area(), s.Perimeter())

	var s1 Shape = Rectangle{4.0, 6.0}

	fmt.Printf("Rectangle \t: Area = %f, Perimeter = %f\n", s1.Area(), s1.Perimeter())

	var s2 Shape = Triangle{5, 7}

	fmt.Printf("Triangle \t: Area = %f, Perimeter = %f\n", s2.Area(), s2.Perimeter())
}
