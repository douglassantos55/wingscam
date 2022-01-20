package main

import (
	"testing"
)

func TestNotEnoughFood(t *testing.T) {
	card := CreateCard("Cockatiel", 3, And(Invertebrate, Rodent), 0, Forest)
	player := CreatePlayer(nil)

	player.DrawCard(card)

	player.GainFood(Fruit)
	player.GainFood(Seed)
	player.GainFood(Rodent)

	if player.PayFood(card.foodCost) {
		t.Errorf("Expected to not have the food cost necessary to play the card")
	}

	if len(player.foods) != 3 {
		t.Errorf("Expected food to have been kept")
	}
}

func TestWithEnoughFood(t *testing.T) {
	card := CreateCard("Cockatiel", 3, And(Invertebrate, Rodent), 0, Forest)
	player := CreatePlayer(nil)

	player.DrawCard(card)

	player.GainFood(Rodent)
	player.GainFood(Fish)
	player.GainFood(Invertebrate)

	if !player.PayFood(card.foodCost) {
		t.Errorf("Expected to have the food cost necessary to play the card")
	}

	if len(player.foods) != 1 {
		t.Errorf("Expected food to have been removed")
	}
}

func TestWithEnoughOrFood(t *testing.T) {
	card := CreateCard("Cockatiel", 3, Or(Invertebrate, Rodent), 0, Forest)
	player := CreatePlayer(nil)

	player.DrawCard(card)

	player.GainFood(Fish)
	player.GainFood(Seed)
	player.GainFood(Fruit)
	player.GainFood(Invertebrate)

	if !player.PayFood(card.foodCost) {
		t.Errorf("Expected to have the food cost necessary to play the card")
	}

	if len(player.foods) != 3 {
		t.Errorf("Expected food to have been removed")
	}
}

func TestNotEnoughOrFood(t *testing.T) {
	card := CreateCard("Cockatiel", 3, Or(Invertebrate, Rodent), 0, Forest)
	player := CreatePlayer(nil)

	player.DrawCard(card)

	player.GainFood(Fruit)
	player.GainFood(Seed)

	if player.PayFood(card.foodCost) {
		t.Errorf("Expected to not have the food cost necessary to play the card")
	}

	if len(player.foods) != 2 {
		t.Errorf("Expected food to have been kept")
	}
}

func TestNotEnoughSingleFood(t *testing.T) {
	card := CreateCard("Cockatiel", 3, Single(Invertebrate), 0, Forest)
	player := CreatePlayer(nil)

	player.DrawCard(card)

	player.GainFood(Fruit)
	player.GainFood(Seed)

	if player.PayFood(card.foodCost) {
		t.Errorf("Expected to not have the food cost necessary to play the card")
	}

	if len(player.foods) != 2 {
		t.Errorf("Expected food to have been kept")
	}
}

func TestEnoughSingleFood(t *testing.T) {
	card := CreateCard("Cockatiel", 3, Single(Invertebrate), 0, Forest)
	player := CreatePlayer(nil)

	player.DrawCard(card)

	player.GainFood(Fruit)
	player.GainFood(Seed)
	player.GainFood(Invertebrate)

	if !player.PayFood(card.foodCost) {
		t.Errorf("Expected to not have the food cost necessary to play the card")
	}

	if len(player.foods) != 2 {
		t.Errorf("Expected food to have been removed")
	}
}

func TestPayEggCost(t *testing.T) {
	card := CreateCard("Cockatiel", 3, Single(Invertebrate), 2, Forest)
	other := CreateCard("Mocking", 3, Single(Invertebrate), 0, Forest)
	player := CreatePlayer(nil)

	player.PlayCard(other)
	other.LayEggs(3)

	if !player.PayEggs(card.eggsCost) {
		t.Errorf("Expected to have enough eggs")
	}

	if other.CountEggs() != 1 {
		t.Errorf("Expected 1 egg to remain, got %d", other.CountEggs())
	}
}

func TestNotEnoughEggs(t *testing.T) {
	card := CreateCard("Cockatiel", 3, Single(Invertebrate), 3, Forest)
	other := CreateCard("Mocking", 3, Single(Invertebrate), 0, Forest)
	player := CreatePlayer(nil)

	player.PlayCard(other)
	other.LayEggs(2)

	if player.PayEggs(card.eggsCost) {
		t.Errorf("Expected to not have enough eggs")
	}

	if other.CountEggs() != 2 {
		t.Errorf("Expected 2 eggs to remain on card, got %d", other.CountEggs())
	}
}

func TestEggsInMultipleCards(t *testing.T) {
	other := CreateCard("Mocking", 3, Single(Seed), 0, Forest)
	another := CreateCard("Mocking", 3, Single(Seed), 0, Forest)
	card := CreateCard("Cockatiel", 3, Single(Invertebrate), 2, Forest)

	player := CreatePlayer(nil)

	other.LayEggs(1)
	another.LayEggs(1)

	player.PlayCard(other)
	player.PlayCard(another)

	if !player.PayEggs(card.eggsCost) {
		t.Errorf("Expected to have enough eggs")
	}

	if other.CountEggs() != 0 || another.CountEggs() != 0 {
		t.Errorf("Expected eggs to be removed from cards, got %d and %d", other.CountEggs(), another.CountEggs())
	}
}

func TestEggsInMultipleCardsNotEnough(t *testing.T) {
	other := CreateCard("Mocking", 3, Single(Seed), 0, Forest)
	another := CreateCard("Mocking", 3, Single(Seed), 0, Forest)
	card := CreateCard("Cockatiel", 3, Single(Invertebrate), 3, Forest)

	player := CreatePlayer(nil)

	other.LayEggs(1)
	another.LayEggs(1)

	player.PlayCard(other)
	player.PlayCard(another)

	if player.PayEggs(card.eggsCost) {
		t.Errorf("Expected to not have enough eggs")
	}

	if other.CountEggs() != 1 || another.CountEggs() != 1 {
		t.Errorf("Expected eggs to be remain, got %d and %d", other.CountEggs(), another.CountEggs())
	}
}

func TestEggsInMultipleCardsOnlyOneThough(t *testing.T) {
	other := CreateCard("Mocking", 3, Single(Seed), 0, Forest)
	another := CreateCard("Mocking", 3, Single(Seed), 0, Forest)
	card := CreateCard("Cockatiel", 3, Single(Invertebrate), 2, Forest)

	player := CreatePlayer(nil)

	other.LayEggs(1)
	another.LayEggs(2)

	player.PlayCard(other)
	player.PlayCard(another)

	if !player.PayEggs(card.eggsCost) {
		t.Errorf("Expected to have enough eggs")
	}

	if other.CountEggs() != 0 || another.CountEggs() != 1 {
		t.Errorf("Expected 0 and 1, got %d and %d", other.CountEggs(), another.CountEggs())
	}
}
