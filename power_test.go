package main

import "testing"

func TestItActivatesWhenPlayedPower(t *testing.T) {
    card := CreateCard("name", 3, nil, 0, Forest)
    card.GivePower(WhenPlayed, LayEggAction{2, card})

    player := CreatePlayer(nil)

    player.PlayCard(card)

    if card.CountEggs() != 2 {
        t.Errorf("Expected card to have 2 eggs, got %d", card.CountEggs())
    }
}

func TestItActivatesPowersFromLeftToRight(t *testing.T) {
    card := CreateCard("name", 6, nil, 0, Forest)
    other := CreateCard("other", 3, nil, 0, Forest)
    another := CreateCard("another", 3, nil, 0, Forest)
    yetanother := CreateCard("yetanother", 3, nil, 0, Forest)

    card.GivePower(WhenActivated, LayEggAction{2, card})
    other.GivePower(WhenActivated, LayEggAction{2, other})
    another.GivePower(WhenPlayed, LayEggAction{2, another})
    yetanother.GivePower(WhenActivated, LayEggAction{2, yetanother})

    player := CreatePlayer(nil)

    player.PlayCard(card)
    player.PlayCard(other)
    player.PlayCard(another)
    player.PlayCard(yetanother)

    if card.CountEggs() != 6 {
        t.Errorf("Expected card to have 6 eggs, got %d", card.CountEggs())
    }

    if other.CountEggs() != 3 {
        t.Errorf("Expected other to have 3 eggs, got %d", other.CountEggs())
    }

    if another.CountEggs() != 2 {
        t.Errorf("Expected another to have 2 eggs, got %d", another.CountEggs())
    }

    if yetanother.CountEggs() != 0 {
        t.Errorf("Expected yetanother to have 0 eggs, got %d", yetanother.CountEggs())
    }
}
