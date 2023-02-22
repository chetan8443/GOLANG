package main

import (
	"fmt"
	
)

func main() {
	fmt.Println("Enter a digit for sum : ")
	var num int
	fmt.Scan(&num)

	var sum,rem int = 0,0     //assigining values in sum,rem variables
	for num!=0{
		rem=num%10                 
		num=num/10	
		sum=sum+rem
		
	}
	fmt.Println("sum is :" ,sum)    //printing sum of digits

	
}
