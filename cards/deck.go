package main

import "fmt"

// CREATE A NEW TYPE OF 'DECK'
// WHICH IS A SLICE OF STRINGS
type deck []string

func(d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}