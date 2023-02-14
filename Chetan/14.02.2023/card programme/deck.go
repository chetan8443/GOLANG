package main

import (
	"fmt"
	"io/ioutil"

	"strings"
)



type deck []string

func (d deck) print() {

	for i, card := range d {
		fmt.Println(i, card)
	}
}

func newDeck() deck {
	cards := deck{}
	cardSuits := deck{"spades ", "hearts ", "diamonds ", "clubs "}
	cardValues := deck{"ace,", "two,", "three,", "four,", "five,", "six,", "seven,", "eight,", "nine,", "joker,", "queen,", "king,"}

	for _, suits := range cardSuits {
		for _, values := range cardValues {
			cards = append(cards, suits+"of "+values)
			
		}
	}
	return cards
}

func deal(d deck,handsize int)(deck,deck){

return d[:handsize],d[handsize:]
}

func (d deck) toString() string{
	return strings.Join(d,",")
}

func (d deck)saveToFile(filename string) error{
return ioutil.WriteFile(filename,[]byte(d.toString()),0666)
}

func newDeckFrimFile(filename string)deck{
 bs,err :=ioutil.ReadFile(filename)
 if err!=nil{
	println("err :",err)
 }
}
