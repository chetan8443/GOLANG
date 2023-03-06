/*
Staircase detail
This is a staircase of size :

		  #
		 ##
		###
	   ####

Sample Input
6
Sample Output

	    #
	   ##
	  ###
	 ####
	#####

######
Explanation
The staircase is right-aligned, composed of # symbols and spaces, and has a height and width of .
*/
package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			fmt.Print(" ")
		}
		for j := 0; j <= i; j++ {
			fmt.Print("#")
		}
		fmt.Println()
	}

}
