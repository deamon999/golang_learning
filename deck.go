package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// create a new type of 'deck
// which is a slice of strings
type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) saveToFile(fileName string) error {
	return os.WriteFile(fileName, []byte(d.toString()), 0666)
	// return ioutil.WriteFile(fileName, []byte(d.toString()), 0666)
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

func readFromFile(fileName string) deck {
	b, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	return deck(strings.Split(string(b), ","))
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
