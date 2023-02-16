package main
import "fmt"

func rev(arr []int)[]int{
	if len(arr)==0{
		return arr
	}
	return append(rev(arr[1:]),arr[0])
}
func main(){
	arrInt:=[]int{1,2,3,4,5}
	fmt.Println(arrInt,rev(arrInt))	
}