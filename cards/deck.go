package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

//create a new type of deck which is a slice of strings

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Hearts", "Clubs", "Diamonds"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value + " of " + suit)
		}
	}

	return cards
}

// this is receiver
func (cards deck) print() {
	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	// return one single string from deck
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	// write to file with permissions 0666
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}