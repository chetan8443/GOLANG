package main

import "fmt"

func main() {

	nums := []int{1, 2, 3, 4, 5}

	p := &nums[0]

	*p = 10

	fmt.Println(nums)
}
