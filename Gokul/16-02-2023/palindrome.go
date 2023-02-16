//Program to check if a number entered is a palindrome or not
package main

import "fmt"

func main() {
	var num int
	fmt.Println("Enter number")
	fmt.Scan(&num)
	fmt.Print(num," is ",palin(num))
}
func palin(n int) string {
	var rem, num_orig int
	num_orig = n
	res :=0
	for n >0{
		rem  = n%10
		res  = res *10 + rem 
		n    =  n/10
	}
	if res == num_orig {
		return "Palindrome"
	}else {
		return "Not Palindrome"
	}
}