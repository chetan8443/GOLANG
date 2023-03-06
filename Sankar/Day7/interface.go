package main
import (
	"fmt"
	"math"
)
func main(){
	c:= Circle{5}
	fmt.Printf("Circle with Radius:= %1.2f , having Area = %1.2f and Perimeter= %1.2f \n",c.Radius,c.Area(),c.Perimeter())
	t:=Triangle{3,4,5}
	fmt.Println("Tringale area and perimeter ","(",t.Area(),",",t.Perimeter(),")")
	r:=Rectangle{2,3}
	fmt.Println("Rectangle")
	fmt.Println("Area:",r.Area())
	fmt.Println("Perimeter:",r.Perimeter())

}
type Circle struct{
	Radius float64
}
type Rectangle struct{
	Length float64
	Width float64
}
type Triangle struct{
	x float64
	y float64
	z float64
}
type Shape interface{
	Area() float64
	Perimeter() float64
}
func (c Circle)Area() float64{
	return c.Radius*c.Radius*math.Pi
}
func (c Circle)Perimeter()float64{
	return 2*math.Pi*c.Radius
} 
func (r Rectangle)Area() float64{
	return r.Width*r.Length
}
func (r Rectangle) Perimeter() float64{
	return 2*r.Width + 2*r.Length
}
func (t Triangle)Area() float64{
	s := (t.x + t.y + t.z) / 2
	return 0.5*(s * (s - t.x) * (s - t.y) * (s - t.z))
}
func (t Triangle) Perimeter() float64{
	return t.x+t.y+t.z
}


