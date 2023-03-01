package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

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

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

type Triangle struct {
	A float64
	B float64
	C float64
}

func (t Triangle) Area() float64 {
	s := (t.A + t.B + t.C) / 2
	return math.Sqrt(s * (s - t.A) * (s - t.B) * (s - t.C))
}

func (t Triangle) Perimeter() float64 {
	return t.A + t.B + t.C
}

func main() {
	r := Rectangle{Width: 3, Height: 4}
	fmt.Println("Rectangle Area:", r.Area())
	fmt.Println("Rectangle Perimeter:", r.Perimeter())

	c := Circle{Radius: 5}
	fmt.Println("Circle Area:", c.Area())
	fmt.Println("Circle Perimeter:", c.Perimeter())

	t := Triangle{A: 3, B: 4, C: 5}
	fmt.Println("Triangle Area:", t.Area())
	fmt.Println("Triangle Perimeter:", t.Perimeter())
}
