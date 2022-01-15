package main

import "testing"

func TestRemoveCardFromDeck(t *testing.T) {
    mocking := &Card{name: "Mocking bird"}
    cockatiel := &Card{name: "Cockatiel"}

    deck := CreateDeck(mocking, cockatiel)
    drawn := deck.Draw()

    if deck.Draw() == drawn {
        t.Error("Should remove card from top of the deck")
    }
}

func TestPlayerReceivesCard(t *testing.T) {
    mocking := &Card{name: "Mocking bird"}
    cockatiel := &Card{name: "Cockatiel"}

    deck := CreateDeck(mocking, cockatiel)
    player := CreatePlayer(deck)

    player.DrawCards(2)

    if player.CountCardsInHand() != 2 {
        t.Error("Player should receive drawn card")
    }
}

func TestDrawFromEmptyDeck(t *testing.T) {
    deck := CreateDeck()
    player := CreatePlayer(deck)

    player.DrawCards(2)

    if player.CountCardsInHand() != 0 {
        t.Error("Player should not get cards from empty deck")
    }
}
