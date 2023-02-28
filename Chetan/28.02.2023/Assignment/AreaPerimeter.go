package main

import "fmt"

type shapes interface {
	area() float64
	Perimeter()	float64
}

type Rectangle struct {
	length float64
	width  float64
}
type Triangle struct {
	base   float64
	height float64
	side1  float64
	side2  float64
}
type Circle struct {
	radius float64
}

func main() {

	rec := Rectangle{length: 32, width: 23}
	tri := Triangle{base: 43, height: 23, side1: 1, side2: 2}
	cir := Circle{radius: 4}

	fmt.Println("Rectangle AREA :", areaPrint(rec))
	fmt.Println("triangle AREA :", areaPrint(tri))
	fmt.Println("Circle AREA :", areaPrint(cir))

	fmt.Println("Rectangle PERIMETER :", perimeterPrint(rec))
	fmt.Println("triangle PERIMETER :", perimeterPrint(tri))
	fmt.Println("Circle PERIMETER :", perimeterPrint(cir))

}

func areaPrint(c shapes) float64 {
	return c.area()
}
func perimeterPrint(c shapes) float64 {
	return c.Perimeter()
}

//All shapes area functions are below
func (t Triangle) area() float64 {

	return 0.5 * t.base * t.height
}
func (r Rectangle) area() float64 {

	return r.width * r.length
}
func (c Circle) area() float64 {

	return 3.14 * c.radius * c.radius
}

//All shapes perimeter functions are below
func (r Rectangle) Perimeter() float64 {

	return 2 * (r.width + r.length)
}
func (t Triangle) Perimeter() float64 {

	return t.side1 + t.side2 + t.base
}

func (c Circle) Perimeter() float64 {

	circumference := 2 * 3.14 * c.radius
	return circumference
}
