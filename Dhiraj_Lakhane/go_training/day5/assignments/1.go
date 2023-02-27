//Write a program that declares an array of integers and prints out the sum of the elements.
package main
 import "fmt"

 func main(){
	// var n int
	// fmt.Scanln(&n)
	rem()
	sortString()
	var arr [5]int
	
	fmt.Printf("Enter array elements: \n")
	for i := 0; i <5; i++ {
		fmt.Printf("Elements: arr[%d]: ", i)
		fmt.Scanln(&arr[i])
	}
	
	 sol(arr)

	 
}
func sol(arr[5]  int){
	var sum int
	for i:=range arr{
		sum+=arr[i]
	}
	fmt.Println(sum)
}
