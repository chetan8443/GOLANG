package main

import "fmt"

func primeNo(num int) {
	flag := true
	if num == 1 {
		fmt.Printf("The %d is not a Prime number\n", num)
	} else {
		for i := 2; i < num; i++ {
			if num%i == 0 {
				flag = false
				break
			}
		}
		if flag == true {
			fmt.Printf("%d is Prime Number\n", num)
		} else {
			fmt.Printf("%d is Not Prime Number\n", num)
		}
	}
}

func pali(num int) {
	rev := 0
	for i := num; i > 0; i /= 10 {
		rem := i % 10
		rev = rev*10 + rem
	}
	if rev == num {
		fmt.Printf("%d is palindrome Number\n", num)
	} else {
		fmt.Printf("%d is Not Palindrome number\n", num)
	}
}

func avg(arr [5]int) float32 {
	var sum int
	var avg float32
	for i := 0; i < len(arr); i++ {
		sum = sum + arr[i]
	}
	avg = float32(sum) / float32(len(arr))
	return avg
}
