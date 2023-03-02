package main

import "fmt"

func main() {

	slice := []int{1, 2, 3, 4, 5}

	fmt.Println(slice)

	reverse(&slice) //function to reverse

}

// takes pointer as arguement
func reverse(r *[]int) {
	slice := *r
	rev := []int{}
	for i := range slice {
		rev = append(rev, slice[len(slice)-1-i])
	}
	fmt.Println(rev)
}
