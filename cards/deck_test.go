package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {

	d := createDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck lenght of 52, but got %v", len(d))
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {

	os.Remove("_decktesting")

	d := createDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck lenght of 52, but got %v", len(d))
	}

	d.saveToFile("_decktesting")

	loadDeck := newDeckFromFile("_decktesting")

	if len(loadDeck) != 52 {
		t.Errorf("Expected deck lenght of 52, but got %v", len(loadDeck))
	}

	os.Remove("_decktesting")
}
