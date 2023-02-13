package main

func main() {
	// cards := deck{}
	// cards := []string{newCard(), newCard(), "Four of Diamonds"}
	// cards := deck{newCard(), newCard(), "Four of Diamonds"}
	// cards = append(cards, "New Card Data")

	// for i, card := range cards {
	// 	fmt.Println(i, card)
	// }

	// cards.print()

	// cards := newDeck()
	// cardsDist, remainingCards := deal(cards, 5)
	// cardsDist.print()
	// remainingCards.print()

	// saveFile(cards)

	// fmt.Println(readFile("saved_cards"))

	cards := newDeck()
	cards.shuffle()
	cards.print()
}
