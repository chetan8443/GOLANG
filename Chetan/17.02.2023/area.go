package main
import "fmt"


type shapes interface{
	area() int

}

			type rectangle struct{
				length int
				height int
			}
			type circle struct{
				radius int
				
			}
			
			func main(){
				
					rec:=rectangle{length: 32,height: 23}
					cir:= circle{radius: 43}
			
			
				fmt.Println("RECTANGLE AREA :",areaPrint(rec))
				fmt.Println("CIRCLE AREA :",areaPrint(cir))
				
			}

			
			func areaPrint(c shapes) int{
				return c.area()
			}
			
	
func (c circle)area() int{
	p:=3*c.radius*c.radius
	return p
}
func (r rectangle)area() int{
	p:=r.height*r.length
	return p
}
