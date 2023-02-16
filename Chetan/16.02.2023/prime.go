package main

import "fmt"

func main() {
	fmt.Print("enter no to check prime or not :")
	var num int
	fmt.Scan(&num)
	flag := false
	for i := 2; i < num; i++ {

		if num != 2 || num != 3 {

			if num%2 == 0 {

				flag = false
				break
			} else if num%i == 0 {

				flag = false
				break
			} else {

				flag = true
			
				

			}
		}
	}
	if flag == true {
		fmt.Print("Number is prime")
	} else {
		fmt.Print(" it is not a Prime ")

	}
}
