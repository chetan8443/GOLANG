package main
import "fmt"

func area(length,breadth int)int {
	ar := length * breadth
	return ar
}
func swap(a,b *int)int{
	var o int
	o=*a
	*a=*b
	*b=o
	return o
}
func main(){
	// fmt.Printf("Area of rectangle is :%d",area(35,25))

	var p int=10
	var q int=20

	fmt.Printf("p=",p,"& q=",q)
	swap(&p,&q)
	fmt.Printf("p=",p,"& q=",q)
}