//Write a program that defines an interface type called Shape with 2 methods: Area() and perimeter(),
//both of which return a float 64 value.
//Implement this interface for 3 diffrent types : Rectangle , Circle , Triangle

package main

import (
	"fmt"
	"math"
)

func main() {
	r :=
		rectangle{
			length:  10,
			breadth: 20,
		}

	c :=
		circle{
			radius: 30,
		}

	t :=
		triangle{
			side:   40,
			height: 50,
		}
	fmt.Println(r)
	fmt.Println(c)
	fmt.Println(t)

	calculate(r)
	calculate(c)
	calculate(t)

}

type rectangle struct {
	length  float64
	breadth float64
}

type circle struct {
	radius float64
}

type triangle struct {
	side   float64
	height float64
}

func (r rectangle) area() float64 {
	return r.length * r.breadth
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (t triangle) area() float64 {
	return (t.height * t.side) / 2
}

func (r rectangle) perimeter() float64 {
	return 2 * r.length * r.breadth
}
func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}
func (t triangle) perimeter() float64 {
	return 3 * t.side
}

type shape interface {
	area() float64
	perimeter() float64
}

func calculate(s shape) {

	fmt.Println("area is :", s.area(), "sq. mt.")
	fmt.Println("perimeter is:", s.perimeter(), "mt.")
}
