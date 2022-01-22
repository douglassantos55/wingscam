package main

import "testing"

func TestGainFood(t *testing.T) {
	player1 := CreatePlayer(nil)
	game := CreateGame(player1)
	game.GainFood()

	if len(player1.foods) != 1 {
		t.Errorf("Expected player1 to have 1 food, got %d", len(player1.foods))
	}
}

func TestGainFoodActivatesCardsPower(t *testing.T) {
	card := CreateCard("name", 2, nil, 0, Forest)
	card.GivePower(WhenActivated, LayEggPower{1, card})

	other := CreateCard("other", 2, nil, 0, Grassland)
	other.GivePower(WhenActivated, LayEggPower{1, other})

	another := CreateCard("another", 2, nil, 0, Forest)
	another.GivePower(WhenActivated, LayEggPower{1, another})

	deck := CreateDeck(another, other, card)
	player1 := CreatePlayer(deck)
	game := CreateGame(player1)

	player1.DrawCards(3)
	game.PlayCard()
	game.PlayCard()
	game.PlayCard()
	game.GainFood()

	if card.CountEggs() != 1 {
		t.Errorf("Expected card to have 1 egg, got %d", card.CountEggs())
	}

	if other.CountEggs() != 0 {
		t.Errorf("Expected other to have 0 egg, got %d", other.CountEggs())
	}

	if another.CountEggs() != 1 {
		t.Errorf("Expected another to have 1 egg, got %d", another.CountEggs())
	}
}

func TestLayEggs(t *testing.T) {
	player1 := CreatePlayer(nil)
	game := CreateGame(player1)

	card := CreateCard("name", 2, nil, 0, Forest)
	player1.DrawCard(card)

	game.PlayCard()
	game.LayEggs()

	if card.CountEggs() != 2 {
		t.Errorf("Expected card to have 2 eggs, got %d", card.CountEggs())
	}
}

func TestLayEggsActivatesCardsPower(t *testing.T) {
	player1 := CreatePlayer(nil)
	game := CreateGame(player1)

	card := CreateCard("name", 2, nil, 0, Grassland)
	card.GivePower(WhenActivated, LayEggPower{1, card})

	other := CreateCard("other", 2, nil, 0, Grassland)
	other.GivePower(WhenActivated, LayEggPower{1, other})

	another := CreateCard("another", 12, nil, 0, Forest)

	player1.DrawCard(another)
	player1.DrawCard(other)
	player1.DrawCard(card)

	game.PlayCard()
	game.PlayCard()
	game.PlayCard()
	game.LayEggs()

	if card.CountEggs() != 1 {
		t.Errorf("Expected card to have 1 eggs, got %d", card.CountEggs())
	}

	if other.CountEggs() != 1 {
		t.Errorf("Expected other to have 1 eggs, got %d", other.CountEggs())
	}

	if another.CountEggs() != 3 {
		t.Errorf("Expected another to have 3 eggs, got %d", another.CountEggs())
	}
}

func TestDrawCards(t *testing.T) {
	card := CreateCard("name", 2, nil, 0, Wetland)
	another := CreateCard("another", 2, nil, 0, Wetland)
	deck := CreateDeck(card, another)
	player1 := CreatePlayer(deck)
	game := CreateGame(player1)

	game.DrawCards()

	if player1.CountCardsInHand() != 1 {
		t.Errorf("Expected player to get 1 card, got %d", player1.CountCardsInHand())
	}
}

func TestDrawCardsActivatesPowers(t *testing.T) {
	card := CreateCard("name", 2, nil, 0, Wetland)
	card.GivePower(WhenActivated, LayEggPower{1, card})

	another := CreateCard("another", 2, nil, 0, Wetland)
	another.GivePower(WhenActivated, LayEggPower{1, another})

	deck := CreateDeck(card, another)
	player1 := CreatePlayer(deck)

	player1.DrawCard(another)
	player1.DrawCard(card)

	game := CreateGame(player1)

	game.PlayCard()
	game.PlayCard()

    // Drawing here gets 2 cards again, cause 3rd column
	game.DrawCards()

	if player1.CountCardsInHand() != 2 {
		t.Errorf("Expected player to have 2 cards, got %d", player1.CountCardsInHand())
	}

	if card.CountEggs() != 1 {
		t.Errorf("Expected card to have 1 eggs, got %d", card.CountEggs())
	}

	if another.CountEggs() != 1 {
		t.Errorf("Expected another to have 1 eggs, got %d", another.CountEggs())
	}
}
