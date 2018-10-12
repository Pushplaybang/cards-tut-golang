package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("expected deck length of 16, but got: %v", len(d))
	}

	if d[0] != "ace of spades" {
		t.Errorf("Expected first card of: ace of spades but instead got %v", d[0])
	}

	if d[len(d) - 1] != "four of clubs" {
		t.Errorf("Expected last card of: four of clubs but instead got %v", d[len(d) - 1])
	}
}

func TestNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in deck loaded from file, but got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}
