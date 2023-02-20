/*A perfect number is a number which is equal to the sum of its divisors(excluding the number itself). This pgm checks whether
a number is perfect or not. Some egs of perfect numbers are 6, 28, 496*/

package main

import "fmt"

func main() {
	fmt.Println("Enter the number")
	var n int
	fmt.Scan(&n)
	perfect(n)
}

func perfect(num int){
	sum:=1
	for i:=2;i*i<=num; i++ {
		if(num % i ==0){
			if( i*i !=num){
				sum = sum + i + num/i
			}else{
				sum = sum + i
			}
		}
	}
	if(sum == num) && num !=1{
		fmt.Println(num,"is a perfect number")
	}else{
		fmt.Println(num,"is not a perfect number")
	}
}
