package main

import (
	"fmt"
	"math/rand"
)

type card struct {
	attack  int
	defence int
}

type deck []card

func (d deck) showDeck() {
	fmt.Println()
	for i, card := range d {
		fmt.Printf("[%d] Attack: [%2d] | Defence: [%2d]\n", i+1, card.attack, card.defence)
	}
}

func newCard() card {
	attack := rand.Intn(15) + 1
	defence := rand.Intn(15) + 1
	newCard := card{attack: attack, defence: defence}
	return newCard
}

func (d *deck) initializeUltimateDeck() {

	for i := 0; i < 50; i++ {
		*d = append(*d, newCard())
	}

}

func (d deck) generateHand(handSize int) (deck, deck) {
	range1 := rand.Intn(50 - handSize)
	range2 := rand.Intn(50 - handSize)
	return d[range1 : range1+handSize], d[range2 : range2+handSize]
}

func (d deck) removeCardByIndex(index int) (deck, card) {

	removed := d[index]
	copy(d[index:], d[index+1:])
	//d[len(d)-1] =
	d = d[:len(d)-1]
	return d, removed
}

func playGame(card1 card, card2 card, gameMode string) int {

	fmt.Printf("\nYou chose: Attack: [%2d] | Defence: [%2d]\n", card1.attack, card1.defence)
	fmt.Printf("Your opponent chose: Attack: [%2d] | Defence: [%2d]\n", card2.attack, card2.defence)

	var value1 int
	var value2 int

	if gameMode == "A" || gameMode == "a" {
		value1 = card1.attack
		value2 = card2.defence
	} else {
		value1 = card1.defence
		value2 = card2.attack
	}

	if value1 > value2 {
		fmt.Println("\nYou won this round!")
		return 1
	} else if value1 == value2 {
		fmt.Println("\nIt's a draw!")
		return 0
	}

	fmt.Println("\nYou lost this round!")
	return -1

}
