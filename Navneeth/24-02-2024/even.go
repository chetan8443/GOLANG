//removes even numbers from slice

package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8}

	for i := 0; i < len(slice); i++ {
		if slice[i]%2 == 0 { //checks if element is even or not
			slice = append(slice[:i], slice[i+1:]...) //removes the element from slice
		}
	}
	fmt.Println(slice) //prints out the new element
}
