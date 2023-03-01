// Implementing Interfaces in Go
package main

import (
	"fmt"
	"math"

)
// declaring an interface type called shape
// an interface contains only the signatures of the methods, but not their implementation
type shape interface{
	area() float64
	perimeter() float64
}

// declaring 3 struct types that represent geometrical shapes: rectangle , circle triangle
type trianle struct{
	base,height float64
	a,b,c float64
}
type rectangle struct{
	width, height float64
}

type circle struct{
	radius float64
}
// method that calculates triangles's area
func (t trianle ) area()float64{
	return (t.base * t.height)/2
	}
// method that calculates triangles's perimeter
func (t trianle) perimeter()float64{
	return (t.a+t.b+t.c)
}
// method that calculates circle's area
func (c circle) area()float64{
return math.Pi * math.Pow(c.radius,2)
}
// method that calculates circle's perimeter
func (c circle) perimeter()float64{
	return 2 * math.Pi * c.radius
}
// method that calculates rectangle's area
func (r rectangle) area()float64{
	return r.height * r.width
	}
// method that calculates rectangle's perimeter
	func (r rectangle) perimeter()float64{
		return 2*(r.width + r.height)
	}

	

	func printShape(s shape)  {
		fmt.Printf("Shape:%T\n ",s)
	 fmt.Println("Area: ",s.area())
	  fmt.Println("Perimeter: ",s.perimeter())
	}

func main() {

	r1:=rectangle{height: 5,width: 6}
	c1:=circle{radius: 7}
	t1:=trianle{a : 2.0, b :3.0, c : 5.0,base: 5,height:10 }
	// printCircle(c1)
	// printRectanle(r1)
	printShape(r1)
	printShape(c1)
	printShape(t1)
}
