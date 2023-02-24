package main

import "fmt"

func main() {

	var x = []int{50, 48, 93, 17, 55, 98, 100, 60}
	var y = []int{} // declaring another slice which is empty
	var a int = len(x)
	for i := 0; i < a; i++ {
		if (x[i] % 2) != 0 {
			y = append(y, x[i])
		}

	}
	fmt.Println("After removing the even number the array is : ", y)

}
