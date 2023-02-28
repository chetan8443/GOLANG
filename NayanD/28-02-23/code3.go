package main

import (
	"fmt"
	"math"
)

// Shape interface with methods for area and perimeter
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Rectangle
type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2*r.Width + 2*r.Height
}

// Circle
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Triangle
type Triangle struct {
	SideA float64
	SideB float64
	SideC float64
}

func (t Triangle) Area() float64 {
	s := (t.SideA + t.SideB + t.SideC) / 2
	return math.Sqrt(s * (s - t.SideA) * (s - t.SideB) * (s - t.SideC))
}

func (t Triangle) Perimeter() float64 {
	return t.SideA + t.SideB + t.SideC
}

func main() {

	r := Rectangle{Width: 5, Height: 7}
	c := Circle{Radius: 4}
	t := Triangle{SideA: 4, SideB: 5, SideC: 6}

	// Print the area and perimeter of each shape
	fmt.Println("Rectangle:")
	fmt.Println("  Area:", r.Area())
	fmt.Println("  Perimeter:", r.Perimeter())

	fmt.Println("Circle:")
	fmt.Println("  Area:", c.Area())
	fmt.Println("  Perimeter:", c.Perimeter())

	fmt.Println("Triangle:")
	fmt.Println("  Area:", t.Area())
	fmt.Println("  Perimeter:", t.Perimeter())
}
