// Write a program to iterate all the elements in the slice

package main

import "fmt"

func main() {
	cards := []string{"Ace of Diamonds", newCard()} // declaration of slice
	cards = append(cards, "six of spades")          // adding new element to slice
	
	for i, card := range cards { 
		fmt.Println(i, card) // printing index and current iteration element
	}
}

func newCard() string {
	return "Five of Diamonds !!"
}
