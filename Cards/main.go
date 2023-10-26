package main

import (
	"fmt"
	"os"
)

var deckSize2 int

func main() {

	cards1 := newDeck2()
	fmt.Println(cards1)
	fmt.Println(cards1.deal2(1))
	fmt.Println(cards1)
	os.Exit(0) ///////////////////////////////////////////////////////////////////////

	cards1.saveToFile("myCards.txt")

	cards1.newDeckFromFile("myCards.txt")

	cards1.print()

	cards1.shuffle()

	cards1.print()

	kards := newDeck()
	//kards.print()

	hand, remainingCards := deal(kards, 5)

	hand.print()
	remainingCards.print()

	/*
	 * golang.org/pkg ==> built-in ve diger kutuphaneler
	 * https://go.dev/play/ ==> go lang playground
	 *
	 * */
	//var card string = "Ace of Spades"
	card := "Adce of Spades"
	// Yukaridaki iki kullanim da ayni anlama geliyor
	// card string olacak ve hep string olarak kalacak
	// anlamina geliyor. card degerine baska bir deger
	// esitleyecegimiz de := degil de = olarak kullaniyoruz
	card = "Ace of Hearts"

	fmt.Println(card)
	//var deckSize1 int
	var deckSize1 = 52
	fmt.Println(deckSize1)
	fmt.Println(deckSize2)

	// Bir variable initialize edilmeden kullanilamaz
	//deckSize3 = 52
	//fmt.Println(deckSize3)

	deckSize4 := newCard()
	fmt.Println(deckSize4)
	cards := []string{"Ace of Diamonds", newCard()}
	fmt.Println("cards: ", cards)

	cards = append(cards, "Six of Spades")
	fmt.Println("cards: ", cards)

	fmt.Println("\n\ncards yazdirilacak...")
	for _, card := range cards {
		fmt.Println(card)
	}
	fmt.Println("\n\ncards yazdirilacak...")
	for i, card := range cards {
		fmt.Println(i, ":", card)
	}

	cardse := deck{"ACE OF DIAMONDS", "KING", "QUEEN"}

	cardse.print()

	bastir("kelimeeee")

}

func bastir(s string) {
	panic("unimplemented")
}

/*
Fonksiyonun bir return degeri var ise fonksiyon
tanimlanirken return type belirtilmesi gerekir
Kendi tanimladigimiz type larda function tanimlama
daha farkli yapiliyor. deck.go yu inceleyebilirsin
*/
func newCard() string {
	return "Five of Diamonds"
}
