//Removes even numbers from slice

package main

import "fmt"

func main() {
	slice := []int{9,10,11,12,13,14,15}

	for i := 0; i < len(slice); i++ {
		if slice[i]%2 == 0 { //checks if element is even or not
			slice = append(slice[:i], slice[i+1:]...) //removes the element from slice
		}
	}
	fmt.Println(slice) //prints out the new element
}