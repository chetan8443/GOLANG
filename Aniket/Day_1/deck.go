package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newCard() string {
	return "Five of Spades"
}

func newDeck() deck {
	cards := deck{}
	newSuits := deck{"Spades", "Dimond", "Black"}
	newValues := deck{"Ace", "Two", "Three", "Four"}
	for _, suit := range newSuits {
		for _, value := range newValues {
			cards = append(cards, value+" of "+suit)

		}
	}
	return cards
}

func deal(d deck, hands int) (deck, deck) {
	return d[:hands], d[hands:]
}

func toString(d deck) string {
	return strings.Join([]string(d), ",")
}

func saveFile(d deck) error {
	return ioutil.WriteFile("saved_cards", []byte(toString(d)), 0666)
}

func readFile(f string) deck {
	bs, err := ioutil.ReadFile(f)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	s := string(bs)
	return strings.Split(s, ",")
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
