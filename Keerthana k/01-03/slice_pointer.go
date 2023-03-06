package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5} //creating slices of interger
	p := &nums[0]                //creating pointer for second element
	*p = 50                      //modify the second element
	fmt.Println(nums)

}
