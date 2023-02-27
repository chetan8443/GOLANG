package main

import "fmt"

func main() {
	arr := [4]int{2, 3, 4,6}          //inititalizing & declaring an array
	var sum int = 0
  
  //executing forloop and sum each element of array 1 by 1
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	fmt.Println(sum)
}
