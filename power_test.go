package main

import "testing"

func TestItActivatesWhenPlayedPower(t *testing.T) {
	card := CreateCard("name", 3, nil, 0, Forest)
	card.GivePower(WhenPlayed, LayEggPower{2, card})

	card.Trigger(WhenPlayed)

	if card.CountEggs() != 2 {
		t.Errorf("Expected card to have 2 eggs, got %d", card.CountEggs())
	}
}
