package main

import "fmt"

// create a new type of 'deck
//which is a slice of strings
type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{" of Spades", " of Diamonds", " of Hearts", " of Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}
	for _, suite := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+suite)
		}
	}
	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
