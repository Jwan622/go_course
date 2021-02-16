package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	testDeck := newDeck()
	deckLength := 16
	if len(testDeck) != deckLength {
		t.Errorf("expected deck legnth of 16 but got %v", len(testDeck))
	}
}

func TestCardsAtTopAndBottom(t *testing.T) {
	testDeck := newDeck()
	if testDeck[0] != "Ace of Spades" {
		t.Errorf("Expected first card Ace of Spades but got %v", testDeck[0])
	}

	if testDeck[len(testDeck) - 1] != "Four of Diamonds" {
		t.Errorf("Expected first card Four of Diamonds but for %v", testDeck[len(testDeck) - 1])
	}
}

func TestSaveToDeckAndNewDeckFromFile (t *testing.T) {
	const fileName = "_decktesting"
	os.Remove(fileName)

	deck := newDeck()
	deck.saveToFile(fileName)

	loadedDeck := newDeckFromFile(fileName)
	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in deck but got %v", len(loadedDeck))
	}

	os.Remove(fileName)
}