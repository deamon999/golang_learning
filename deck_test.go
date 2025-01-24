package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected decl length of 16, but %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first error 'Ace of Spades', but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected first error 'Four of Clubs', but got %v", d[len(d)-1])
	}
}

func TestSaveToFileAndReadFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := readFromFile("_decktesting")
	if len(loadedDeck) != 16 {
		t.Errorf("Expecting 16 card in deck, but got %v", len(deck))
	}

	os.Remove("_decktesting")
}
