package main

import "fmt"

var test int

//var Test int

const test1 = 6

func main() {
	fmt.Println("do the practice")
	test = 5
	fmt.Println("variable of test value is: ", test)
	//test1 = 10
	fmt.Println("varible of test2 value is: ", test1)

	for i := 1; i <= 5; i++ {
		if i == 4 {
			break
		}
		if i == 2 {
			continue
		}
		fmt.Print("\n value of i is: ", i)

	}

}
