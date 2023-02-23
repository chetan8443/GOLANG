package main
import "fmt"
func main() {
	var count int = 4

	for j := 0; j < count; j++ {
  
		for i := 0; i <= j; i++{
			fmt.Print(" ")
		}
		for i := count; i>j; i--{                 
			fmt.Print("*")
		}
		
		fmt.Println("")         //for print to the next line
	}

}
//output
//  ****
//   ***
//    **
//     *
