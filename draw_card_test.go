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

func TestDrawFromTray(t *testing.T) {
    mocking := &Card{name: "Mocking bird"}
    cockatiel := &Card{name: "Cockatiel"}
    goose := &Card{name: "Goose"}

    deck := CreateDeck(mocking, cockatiel, goose)

    tray := CreateTray(3)
    tray.Replenish(deck)

    player := CreatePlayer(deck)
    player.DrawCard(tray.GetCard(0))
    player.DrawCard(tray.GetCard(0))

    if player.CountCardsInHand() != 2 {
        t.Errorf("Expected player to have two cards in hand, got %d", player.CountCardsInHand())
    }

    if tray.Count() != 1 {
        t.Errorf("Expected tray to have one card, got %d", len(tray.cards))
    }

    if deck.Draw() != nil {
        t.Errorf("Expected deck to be empty")
    }
}

func TestGetCardOutOfBoundsFromTray(t *testing.T) {
    mocking := &Card{name: "Mocking bird"}
    cockatiel := &Card{name: "Cockatiel"}
    goose := &Card{name: "Goose"}

    deck := CreateDeck(mocking, cockatiel, goose)

    tray := CreateTray(3)
    tray.Replenish(deck)

    player := CreatePlayer(deck)
    player.DrawCard(tray.GetCard(5))

    if player.CountCardsInHand() != 0 {
        t.Errorf("Expected player to not have cards in hand, got %d", player.CountCardsInHand())
    }
}
