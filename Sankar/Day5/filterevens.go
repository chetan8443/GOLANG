package main 
import "fmt"
func main(){
	array1 := [...]int {1,2,3,4,5,6,6,56,7}
	slice1 := array1 [:]
	result := make([]int,0,18)

	for i:=0;i<len(slice1);i++{
		if (slice1[i])%2==0{
			result = append(result, slice1[i])		
		}
	}
	fmt.Println(result)

}
