package main

import "testing"

// food cost === discards from player's stack, not birds
// habitat
// egg cost === discards from birds (column - board)

func TestNotEnoughFood(t *testing.T) {
    card := CreateCard("Cockatiel", 3, And(Invertebrate, Rodent))
    player := CreatePlayer(nil)

    player.DrawCard(card)

    player.GainFood(Fruit)
    player.GainFood(Seed)
    player.GainFood(Rodent)

    if player.Has(card.foodCost) {
        t.Errorf("Expected to not have the food cost necessary to play the card")
    }
}

func TestWithEnoughFood(t *testing.T) {
    card := CreateCard("Cockatiel", 3, And(Invertebrate, Rodent))
    player := CreatePlayer(nil)

    player.DrawCard(card)

    player.GainFood(Rodent)
    player.GainFood(Fish)
    player.GainFood(Invertebrate)

    if !player.Has(card.foodCost) {
        t.Errorf("Expected to have the food cost necessary to play the card")
    }
}

func TestWithEnoughOrFood(t *testing.T) {
    card := CreateCard("Cockatiel", 3, Or(Invertebrate, Rodent))
    player := CreatePlayer(nil)

    player.DrawCard(card)

    player.GainFood(Fish)
    player.GainFood(Seed)
    player.GainFood(Fruit)
    player.GainFood(Invertebrate)

    if !player.Has(card.foodCost) {
        t.Errorf("Expected to have the food cost necessary to play the card")
    }
}

func TestNotEnoughOrFood(t *testing.T) {
    card := CreateCard("Cockatiel", 3, Or(Invertebrate, Rodent))
    player := CreatePlayer(nil)

    player.DrawCard(card)

    player.GainFood(Fruit)
    player.GainFood(Seed)

    if player.Has(card.foodCost) {
        t.Errorf("Expected to not have the food cost necessary to play the card")
    }
}

func TestNotEnoughSingleFood(t *testing.T) {
    card := CreateCard("Cockatiel", 3, Single(Invertebrate))
    player := CreatePlayer(nil)

    player.DrawCard(card)

    player.GainFood(Fruit)
    player.GainFood(Seed)

    if player.Has(card.foodCost) {
        t.Errorf("Expected to not have the food cost necessary to play the card")
    }
}

func TestEnoughSingleFood(t *testing.T) {
    card := CreateCard("Cockatiel", 3, Single(Invertebrate))
    player := CreatePlayer(nil)

    player.DrawCard(card)

    player.GainFood(Fruit)
    player.GainFood(Seed)
    player.GainFood(Invertebrate)

    if !player.Has(card.foodCost) {
        t.Errorf("Expected to not have the food cost necessary to play the card")
    }
}
