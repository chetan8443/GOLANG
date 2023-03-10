// Program to generate random integers within the given range and prints the prime numbers in it.

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var Range int
	var size int
	fmt.Println("Enter the range :")
	fmt.Scanln(&Range)
	fmt.Println("Enter the number of random numbers you need:")
	fmt.Scanln(&size)
	var arr = make([]int, 0)
	for i := 0; i < size; i++ {
		arr = append(arr, rand.Intn(Range))
	}
	fmt.Println("Array numbers:", arr)
	//fmt.Println(isPrime(arr))
	fmt.Println("Prime numbers are:", isPrime(arr))
}

func isPrime(arr []int) []int {
	var arr1 = make([]int, 0)
	flag := 0
	for i := 0; i <= len(arr)-1; i++ {
		flag = 0
		for j := 2; j < arr[i]/2; j++ {
			if arr[i]%j == 0 {
				//fmt.Println(num)
				flag = 1
				break
				//fmt.Println(arr1)
			}
		}
		if flag == 0 && arr[i] > 1 {
			arr1 = append(arr1, arr[i])
		}
	}
	return arr1
}
