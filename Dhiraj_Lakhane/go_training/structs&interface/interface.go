package main

import "fmt"

type Shape interface{
	area() float64
	perimeter() float64
}
type rect struct{
	length float64
	breadth float64
}
type triangle struct{
	base float64
	height float64
}
type circle struct{
	radius float64
}
func (c circle)area()float64{
	return 3.14 * c.radius * c.radius
}
func (r rect)area()float64{
	return r.length * r.breadth
}
func (r rect)perimeter()float64{
	return 2*(r.length + r.breadth)
}
func (t triangle)area()float64{
	return 1/2 * t.base * t.height
}
func(t triangle)perimeter()float64{
	return 2*t.base+t.height
}

func main(){
	r1:=rect{
		15.5,
		20.5,
	}
	c1:=circle{
		7.2,
	}
	t1:=triangle{
		12.5,
		5.25,
	}
	fmt.Println(c1.area())
	fmt.Println(r1.area())
	fmt.Println(r1.perimeter())
	fmt.Println(r1.area())
	fmt.Println(t1.area())
	fmt.Println(t1.perimeter())

}