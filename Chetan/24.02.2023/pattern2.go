
package main

import "fmt"

func main() {
	var count int = 6

	for j := 1; j < count; j++ {

		for i := count; i >j; i-- {
			fmt.Print(" ")
		}
		for k := 0; k < j; k++ {
			fmt.Print("* ")
		}
		fmt.Println("")
	}
	for j := 1; j < count; j++ {

		
		for k := 0; k < j; k++ {
			fmt.Print(" ")
		}
		for i := count; i >j; i-- {
			fmt.Print("* ")
		}
			fmt.Println("")
	}

}


//      * 
//     * *
//    * * *
//   * * * *
//  * * * * *
//  * * * * *
//   * * * *
//    * * *
//     * *
//      *
