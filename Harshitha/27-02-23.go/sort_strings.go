package main
  
import (
    "fmt"
    "sort"
)
  
// Main function
func main() {
  
    // Creating and initializing slices
    // Using shorthand declaration
    arr := []string{"teja","sushmitha","Harshitha","faiza","dhiraj","Bhavana"}
    // Displaying slices
    fmt.Println(arr)
    // Sorting the slice of strings
    // Using Strings function
    sort.Strings(arr)
  
    // Displaying the result
    fmt.Println("\nSlices(After):")
    fmt.Println("Slice 1 : ", arr)
	//This sorting is done through first preference to capitals and then to every other letters
}