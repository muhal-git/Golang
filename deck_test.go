package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 56 {
		t.Errorf("Expected deck length of 56, but got %v", len(d))
	}
}
