
// Write a program that declares a map of strings to integers and prints out the values in reverse order.

package main

import "fmt"

func main() {

	marks := map[string]int{
		"Jack":  98,
		"James": 81,
		"Jim":   88,
		"Jerry": 79,
	}

	fmt.Println("\nThe elements in map before : ")
	list := []string{}
	for k, v := range marks {
		list = append(list, k)
		fmt.Println(k, v)

	}

	fmt.Println("\nThe elements in map after reverse : ")
	for i := len(list) - 1; i >= 0; i-- {
		fmt.Println(list[i], marks[list[i]])
	}
}
