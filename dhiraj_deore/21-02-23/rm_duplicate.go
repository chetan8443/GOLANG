//write a program to remove duplicate values from slice

package main

import "fmt"

func main() {
	list := []int{1, 4, 2, 6, 7, 4, 5, 7, 8, 5, 3, 5, 7, 3, 2} //slice having duplicate values

	newlist := []int{} // new slice to store unique values

	for i := 0; i < len(list); i++ { //  iterating list of vlaues

		c := 1
		for j := 0; j < len(newlist); j++ { // checking for unique values
			if list[i] == newlist[j] {
				c++ // counting the presence of value
			}
		}

		if c == 1 {
			newlist = append(newlist, list[i]) // appending the value to new list
		}
	}
	fmt.Println(newlist)
}
