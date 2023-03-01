/*
Given two numbers a and b. In one operation you can pick any non negative
integer x and either of a or b. Now if you picked a then replace a with
a&x else if you picked b then replace b with b&x.
Return the minimum number of operation required to make a and b equal.
*/
/*
Input1:
5
12
Output1:The number of required operations are 2

Input2:
100
100
Output2:The required number of operation is 0
*/

package main

import "fmt"

func main() {
	fmt.Print("Enter a value:")
	var a int
	fmt.Scanln(&a)
	fmt.Print("Enter b value:")
	var b int
	fmt.Scanln(&b)
	c1 := 0
	c2 := 0
	if a == b {
		fmt.Print("The required number of operation is 0")
	} else {
		for i := 0; i < 64; i++ { //the size of integer datatype
			if (a>>i)&1 == 0 && (b>>i)&1 == 1 {
				c1 += 1
			} else if (a>>i)&1 == 1 && (b>>i)&1 == 0 {
				c2 += 1
			}
		}
	}
	if c1 == 1 && c2 == 1 {
		fmt.Println("The number of required operations are 2")
	} else {
		fmt.Println("The number of required operations are 1")
	}
}
