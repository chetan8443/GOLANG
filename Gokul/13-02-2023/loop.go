package main

//for understanding loops and slices
import "fmt"
func main(){
	slice := []string{"I am 1st element ","I am 2nd element of a slice"}
	slice = append(slice, "Third element")

	fmt.Println("----Slice elements are----")
	for i,value := range slice{
		fmt.Println(i,value)
	}

	fmt.Println("----Slice elements without index----")
	for _,value := range slice{
		fmt.Println(value)
	}
}