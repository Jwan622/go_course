package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
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

func newDeckFromFile(filename string) deck {
	// no receiver since we don't have a deck. The reeturn type is a deck
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("an error while reading occurred")
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	file_array := strings.Split(string(bs), ",") // converts bites to string
	return deck(file_array) //type convert to deck which extends strings anyway. /we do this so wn use functions that have a deck receiver.
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for index := range d {
		// randomly generate beetween 0 and len of slice - 1
		newPosition := r.Intn(len(d) - 1)

		d[index], d[newPosition] = d[newPosition], d[index]
	}
}