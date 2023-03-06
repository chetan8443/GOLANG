package main

import "fmt"

func main() {
	var a [4][4]int
	fmt.Println("Enter the elements into the array")
	var i, j int = 0, 0
	for i = 0; i < 4; i++ {
		for j = 0; j < 4; j++ {
			fmt.Scanln(&a[i][j])
		}
	}
	fmt.Println(a)
	//left diagnol sum-right diagnol sum
	var lsum int = 0
	var rsum int = 0
	for i = 0; i < 4; i++ {
		for j = 0; j < 4; j++ {
			if i == j {
				lsum = lsum + a[i][j]
			} else if i+j == 3 {
				rsum = rsum + a[i][j]
			} else {
				continue
			}
		}
	}
	fmt.Println(lsum, rsum)
	fmt.Println("The diff b/w left and right diagnol is", lsum-rsum)
}
