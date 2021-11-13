package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

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

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()),0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
	// OPTION #1 -> LOG THE ERROR AND RETURN A CALL TO newDeck()
	// OPTION #2 -> LOG THE ERROR AND ENTIRELY QUIT THE PROGRAM	
	fmt.Println("Error:", err)
	os.Exit(1)
	}

	s := strings.Split(string(bs), ",") 
	return deck(s)
	
}