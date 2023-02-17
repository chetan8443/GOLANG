package main

// import "fmt"

// import "fmt"

func main() {

	cards := newDeck()
	// var hand,reamaining=deal(cards,4)
	// hand.print()
	// reamaining.print()

	// fmt.Printf("cards.toString(): %v\n", cards.toString())

	cards.saveToFile("my_cards")

}

// func new1Card() string {
// 	return "ace of space"
// }
