package main

import (
	"fmt"
	"os"
	"testing"
)

func TestBegin(t *testing.T) {
	fmt.Println("--------------------------------------------------")
	fmt.Println("TEST STARTED")
	fmt.Println("--------------------------------------------------")
}

func TestNewDeck1(t *testing.T) {
	d := newDeck()

	if len(d) != 56 {
		t.Errorf("Expected deck length of 56, but got %v", len(d))
	} else {
		fmt.Println("TestNewDeck1 PASSED")
	}
}

func TestNewDeck2(t *testing.T) {
	d := newDeck()

	if len(d) != 56 {
		t.Errorf("Expected deck length of 56, but got %v", len(d))
	} else {
		fmt.Println("TestNewDeck2 PASSED")
	}
}

func TestFirstAndLastCards(t *testing.T) {
	d := newDeck()

	if d[0] != "Spades of One" {
		t.Errorf("Expected Spades of One, but got %v", d[0])
	} else {
		fmt.Println("TestFirstCard PASSED")
	}

	if d[len(d)-1] != "Cubes of A" {
		t.Errorf("Expected Cubes of A, but got %v", d[len(d)-1])
	} else {
		fmt.Println("TestLastCard PASSED")
	}
}

func TestSaveToDeckAndNewDeskFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeck().newDeckFromFile("_decktesting")

	if len(loadedDeck) != 56 {
		t.Errorf("Expected 56 cars in deck, got %v", len(loadedDeck))
	} else {
		fmt.Println("Save to file and read from file PASSED")
	}

	os.Remove("_decktesting")
}

func TestEnd(t *testing.T) {
	fmt.Println("--------------------------------------------------")
	fmt.Println("TEST FINISHED")
	fmt.Println("--------------------------------------------------")
}
