/*
//Go program to check the given number is Armstrong number or not.
package main

import "fmt"

func main() {
	var num, tempnum, rem int
	var result int = 0
	fmt.Println("Enter a three digit number:=")
	fmt.Scanln(&num)
	tempnum = num
	for {
		rem = tempnum % 10
		result += rem * rem * rem
		tempnum /= 10
		if tempnum == 0 {
			break
		}
	}
	if result == num {
		fmt.Printf("%d is an Armstrong number", num)
	} else {
		fmt.Printf("%d is not an Armstrong number", num)
	}
}
*/

/*
// Go program to check the given number is Armstrong number or not , using Function.
package main

import "fmt"

func armstrong(num int) int {
	temp := 0
	rem := 0
	var result int = 0
	temp = num
	for {
		rem = temp % 10
		result += rem * rem * rem * rem
		temp /= 10
		if temp == 0 {
			break
		}
	}
	if result == num {
		fmt.Printf("%d is an Armstrong number", num)
	} else {
		fmt.Printf("%d is not an Armstrong number", num)
	}
	return num
}
func main() {
	var n int
	fmt.Println("Enter the number:")
	fmt.Scan(&n)
	fmt.Println(armstrong(n))
}
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	var arm int
	fmt.Println("Enter number")
	fmt.Scanln(&arm)
	n := arm
	number := arm
	count := 0
	for arm > 0 {
		arm = arm / 10
		count = count + 1
	}
	fmt.Println(count)
	rem := 0
	sum := 0
	var c int = count
	for n > 0 {
		rem = n % 10
		sum = sum + int(math.Pow(float64(rem), float64(c)))
		n = n / 10
	}
	if sum == number {
		fmt.Println("The given number is an armstron number")
	} else {
		fmt.Println("Not an armstrong number")
	}

}
