package main

import (
	"os"
	"testing"
)

const testFilename = "_decktesting"

func TestNewDeck(t *testing.T) {
	deck := newDeck()
	if len(deck) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(deck))
	}

	if "Ace of Diamonds" != deck[0] {
		t.Errorf("Expected first card of Ace of Diamonds, but got %v", deck[0])
	}

	if "Four of Clubs" != deck[len(deck)-1] {
		t.Errorf("Expected last card of Four of Clubs, but got %v", deck[len(deck)-1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove(testFilename)

	deck := newDeck()
	deck.saveToFile(testFilename)

	loaded := newDeckFromFile(testFilename)

	if len(loaded) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(loaded))
	}
}

func TestShuffle(t *testing.T) {
	deck := newDeck()
	f := deck[0]
	deck.shuffle()

	if f == deck[0] {
		t.Errorf("Expected different cards in first position after shuffle the deck. [%v] == [%v]", f, deck[0])
	}
}
