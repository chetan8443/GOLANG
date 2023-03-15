package main

import "fmt"

type shape interface {
	area() float64
	perimeter() float64
}

type square struct {
	side float64
}
type circle struct {
	radius float64
}

func (s square) perimeter() {
	fmt.Println("The perimeter of the square is:", 2*s.side)
}

func (s square) area() {
	fmt.Println("The area of the square is:", s.side*s.side)
}

func (c circle) perimeter() {
	fmt.Println("The perimeter of the circle is:", 2*3.14*c.radius*c.radius)
}

func (c circle) area() {
	fmt.Println("The area of the circle is:", c.radius*c.radius)
}

func printShape(s shape) {
	fmt.Println(s.area())
	fmt.Println(s.perimeter())
}

func main() {
	shapes := []shape{
		square{length: 15.2},
		circle{radius: 3.4},
	}
	for _, v := range shapes() {
		printShape(v)
	}
}
