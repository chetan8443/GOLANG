// Write a program that declares a slice of strings and sorts it in alphabetical order.
package main

import (
	"fmt"
	"sort"
)

func main() {
	animals := []string{"Rabbit", "Fish", "Dog", "Parrot", "Cat", "Hamster"}
	fmt.Println("Before sorting =", animals)
	fmt.Println("length of animals=", len(animals))
	fmt.Println("length of animals=", cap(animals))
	sort.Strings(animals)
	fmt.Println("After sorting =", animals)
}
