package main

import "fmt"

func main() {
	//card := newCard()
	//cards := deck{newCard(), newCard(), "King of Spades"}
	//fmt.Println(card)
	//fmt.Println(cards)
	//
	//fmt.Println("About to iterate")
	//
	////this is how we iterate
	//for i, card := range cards {
	//	fmt.Println(i, card)
	//}
	cards := newDeck()

	hand, remainingCards := deal(cards, 5)
	// this is how we use a receiver
	hand.print()
	remainingCards.print()
	fmt.Println(cards.toString())

	cards.saveToFile("my_cards")

	new_cards := newDeckFromFile("my_cards")
	fmt.Println("about to print")
	new_cards.print()

	deck_to_shuffle := newDeckFromFile("my_cards")
	deck_to_shuffle.shuffle()
	fmt.Println("about to print shuffled deck...")
	deck_to_shuffle.print()
}

func newCard() string {
	return "Five of Diamonds"
}