// program to check if given key is Present or not:

package main

import "fmt"

func main() {

	age := map[string]int{
		"Meer":  23,
		"Fahad": 20,
		"Ali":   21,
		"Syed":  24,
	}

	isKeyPresent(age)

}

func isKeyPresent(m map[string]int) {

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
