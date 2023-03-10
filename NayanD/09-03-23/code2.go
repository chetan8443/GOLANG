// Write a program to swap values using pointers:

package main

import "fmt"

func main() {
	apple := "red"
	mango := "yellow"

	fmt.Printf("Before swapping \napple : %v\nmango : %v\n", apple, mango)
	swap(&apple, &mango)
	fmt.Printf("\nAfter swapping \napple : %v\nmango : %v\n", apple, mango)

}

func swap(a *string, b *string) {
	*a, *b = *b, *a
}
