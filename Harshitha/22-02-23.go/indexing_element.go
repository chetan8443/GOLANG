package main

import "fmt"

func main() {
	var slice1[5] int
	fmt.Print("Give numbers to make a slice :")
	var num int

	//looping to take inputs for an array
	for i := 0; i < len(slice1); i++ {
		fmt.Scan(&num)
		slice1[i] = num
	}

	//Printing array
	for i, s := range slice1 {
		fmt.Println(i, s)
	}

	fmt.Print("Type number which want to search : ")
	var num2 int
	fmt.Scan(&num2)

	//Using for loop to go through each number and checking if the given number exists or not
	var flag int=0
	for i := 0; i < len(slice1); i++ {
		if slice1[i] == num2 {
			fmt.Println("given number",num2,"is found in",i,"th index of the array")
			flag=1
		}else{
			continue
		}
	}
	if flag==0{
		fmt.Println("Element not found")
	}

}