//Write a program that declares an array of floats and calculates the average of the elements.
package main
 import "fmt"

 func main(){
	// var n int
	// fmt.Scanln(&n)
	var arr1 [5]float64
	
	fmt.Printf("Enter array elements: \n")
	for i := 0; i <5; i++ {
		fmt.Printf("Elements: arr[%d]: ", i)
		fmt.Scanln(&arr1[i])
	}
	 sol1(arr1)
}
func sol1(arr[5]  float64){
	var sum float64
	for i:=range arr{
		sum+=arr[i]
	}
	fmt.Println(sum)
}