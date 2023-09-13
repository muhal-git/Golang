package main

import "fmt"

// create new type "deck"
// which is slice of strings
type deck []string


func newDeck() deck{
	cards := deck{}

	cardSuits := []string{"Spades","Diamonds","Hearts","Cubes"}
	cardValues := []string{"One","Two","Three","Four","Five","Six","Se ven","Eight","Nine","Ten","J","Q","K","A"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, suit + " of " + value)
		}
	}

	return cards
}

// Buradaki d this ve self gibi dusunulebilir
// d receiver olarak adlandirilir
func (d deck) print() {
	fmt.Println("print func of deck is running...")
	for i, card := range d {
		fmt.Println("index:", i, " element:", card)
	}
}

func deal(d deck, handSize int) (deck,deck) {
	
	return d[:handSize], d[handSize:]
}