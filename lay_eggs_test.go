package main

import "testing"

func TestBirdLaysEggs(t *testing.T) {
	card := &Card{eggsLimit: 4}
	card.LayEggs(2)

	if card.CountEggs() != 2 {
		t.Errorf("Expected card to have 2 eggs, but got %d", card.CountEggs())
	}
}

func TestEggsLimitIsRespected(t *testing.T) {
	card := &Card{eggsLimit: 2}
	card.LayEggs(4)

	if card.CountEggs() != 2 {
		t.Errorf("Expected card to have 2 eggs, but got %d", card.CountEggs())
	}
}
