package main 
import "fmt"
func main(){
	array1 := [3] int {12,3,4}
	var sumofarray  int = 0 
	for i:=0;i < len(array1);i++{
		sumofarray = sumofarray + array1[i]		
	}
	fmt.Println("Sum of Element of Array : ",sumofarray)
}
