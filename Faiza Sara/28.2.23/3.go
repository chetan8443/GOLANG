/* program that defines an interface type called Shape with two methods: Area()
and Perimeter(), both of which return a float64 value. Implement this interface for three
different types: Rectangle, Circle, and Triangle */

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

// creating Function to calculate Perimeter of Rectangle
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

// creating Function to calculate Perimeter of Circle
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

// creating Function to calculate Perimeter of Triangle
func (t Triangle) Perimeter() float64 {
	side := math.Sqrt((t.base/2)*(t.base/2) + t.height*t.height)
	return 2*side + t.base
}

func main() {
	var s1 Shape = Circle{10.0}

	fmt.Printf("Circle : Area = %f, Perimeter = %f\n", s1.Area(), s1.Perimeter())

	var s2 Shape = Rectangle{10.0, 5.0}

	fmt.Printf("Rectangle : Area = %f, Perimeter = %f\n", s2.Area(), s2.Perimeter())

	var s3 Shape = Triangle{6, 9}

	fmt.Printf("Triangle : Area = %f, Perimeter = %f\n", s3.Area(), s3.Perimeter())
}
