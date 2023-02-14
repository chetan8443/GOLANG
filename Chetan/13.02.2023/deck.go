package main
import "fmt"



type deck []string

func (d deck) print() {

	for i, card := range d {
		fmt.Println(i, card)
	}
}

func newDeck() deck {
	cards := deck{}
	cardSuits := deck{"spades", "hearts", "diamonds", "clubs"}
	cardValues := deck{"ace", "two", "three", "four", "five", "six", "seven", "eight", "nine", "joker", "queen", "king"}

	for _, suits := range cardSuits {
		for _, values := range cardValues {
			cards = append(cards, suits+"of"+values)
			
		}
	}
	return cards
}
