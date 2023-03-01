package main

import "fmt"

func main(){
	var year int 
	fmt.Println("enter the user input year")
	fmt.Scanln(&year)
	var char string
	if (year%2 == 0 && year%100 != 0 || year%400 == 0){
		char:="yes"
	}else{
		char:="No"
	}
	switch char{
	case char=="yes":
		fmt.Println("Its a leap year")
	case char=="no":
		fmt.Println("Its not a leap year")
	}

}