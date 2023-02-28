package main 
import "fmt"
func main(){
	array2 := [...] float64 {4.0,25.0,45.0,31.0}
	var count float64 = 0
	var avg float64 = 0
	for i:=0;i < len(array2);i++{
		count += array2[i]
		avg = (count/float64(len(array2)))
	}
	fmt.Println(avg)
}
