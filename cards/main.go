package main

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

	// this is how we use a receiver
	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}