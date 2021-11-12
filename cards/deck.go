package main

import "fmt"

// CREATE A NEW TYPE OF 'DECK'
// WHICH IS A SLICE OF STRINGS
type deck []string

func newDeck()  deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func(d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// THIS FUNCTION RETURNS 2 VALUES (DECK, DECK)
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}