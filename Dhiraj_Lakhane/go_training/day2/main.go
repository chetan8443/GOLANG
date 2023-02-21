package main

import (
	"fmt"
	a"new/jumps"
)

func main(){
	arr:=[]int{1,3,5,8,9,2,6,7,6,8,9}
	fmt.Println("Given Array: ",arr)
	/*
	minimum no of jumps to reach an end of array
	 Jump from 1st element to 2nd element as there is only 1 step,
	 now there are three options 5, 8 or 9. If 8 or 9 is chosen then the
	  end node 9 can be reached. So 3 jumps are made.

	  3 (1-> 3 -> 9 -> 9)
	  */
	  fmt.Println("min no of jumps to reach end of an array: ",a.Sol(arr))



}
