// write a program to check if given key is available or not:

package main

import "fmt"

func main() {

	marks := map[string]int{
		"Jack":  98,
		"James": 81,
		"Jim":   88,
		"Jerry": 79,
	}

	checkKey(marks)

}

func checkKey(m map[string]int) {

	var key string
	fmt.Println("Enter a key :")
	fmt.Scan(&key)

	k, ok := m[key]

	if ok {
		fmt.Println("Given key is present with the value :", k)
	} else {
		fmt.Println("Given key is not present ")
	}
	fmt.Println()
}
