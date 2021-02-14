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
}

func newCard() string {
	return "Five of Diamonds"
}