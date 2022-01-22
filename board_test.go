package main

import "testing"

func TestItPlacesCardInProperHabitat(t *testing.T) {
    board := CreateBoard(
        CreateRow(Forest, 5, GainFood),
        CreateRow(Grassland, 5, LayEggs),
        CreateRow(Wetland, 5, DrawCards),
    )

    card := CreateCard("Cockatiel", 5, nil, 0, Forest)
    index := board.PlaceCard(card)

    if index != 0 {
        t.Errorf("Expected column 1 to have a card, got %d", index)
    }

    if board.GetHabitat(Forest).CountCards() != 1 {
        t.Errorf("Expectd forest to have 1 card, got %d", board.GetHabitat(Forest).CountCards())
    }

    if board.GetHabitat(Wetland).CountCards() != 0 {
        t.Errorf("Expectd wetland to have 0 card, got %d", board.GetHabitat(Wetland).CountCards())
    }

    if board.GetHabitat(Grassland).CountCards() != 0 {
        t.Errorf("Expectd grassland to have 0 card, got %d", board.GetHabitat(Grassland).CountCards())
    }
}

func TestItPlacesCardInLeftMostExposedColumn(t *testing.T) {
    board := CreateBoard(
        CreateRow(Forest, 5, GainFood),
        CreateRow(Grassland, 5, LayEggs),
        CreateRow(Wetland, 5, DrawCards),
    )

    card := CreateCard("Cockatiel", 5, nil, 0, Forest)
    other := CreateCard("Mocking", 5, nil, 0, Forest)

    index := board.PlaceCard(card)
    otherIndex := board.PlaceCard(other)

    if index != 0 {
        t.Errorf("Expected column 1 to have a card, got %d", index)
    }

    if otherIndex != 1 {
        t.Errorf("Expected column 2 to have a card, got %d", otherIndex)
    }

    if board.GetHabitat(Forest).CountCards() != 2 {
        t.Errorf("Expectd forest to have 2 card, got %d", board.GetHabitat(Forest).CountCards())
    }

    if board.GetHabitat(Wetland).CountCards() != 0 {
        t.Errorf("Expectd wetland to have 0 card, got %d", board.GetHabitat(Wetland).CountCards())
    }

    if board.GetHabitat(Grassland).CountCards() != 0 {
        t.Errorf("Expectd grassland to have 0 card, got %d", board.GetHabitat(Grassland).CountCards())
    }
}

func TestItPlacesCardsInMultipleHabitats(t *testing.T) {
    board := CreateBoard(
        CreateRow(Forest, 5, GainFood),
        CreateRow(Grassland, 5, LayEggs),
        CreateRow(Wetland, 5, DrawCards),
    )

    card := CreateCard("Cockatiel", 5, nil, 0, Forest)
    other := CreateCard("Mocking", 5, nil, 0, Grassland)
    yetanother := CreateCard("Goose", 5, nil, 0, Wetland)

    index := board.PlaceCard(card)
    otherIndex := board.PlaceCard(other)
    yetAnotherIndex := board.PlaceCard(yetanother)

    if index != 0 {
        t.Errorf("Expected column 1 to have a card, got %d", index)
    }

    if otherIndex != 0 {
        t.Errorf("Expected column 1 to have a card, got %d", otherIndex)
    }

    if yetAnotherIndex != 0 {
        t.Errorf("Expected column 1 to have a card, got %d", yetAnotherIndex)
    }

    if board.GetHabitat(Forest).CountCards() != 1 {
        t.Errorf("Expectd forest to have 1 card, got %d", board.GetHabitat(Forest).CountCards())
    }

    if board.GetHabitat(Wetland).CountCards() != 1 {
        t.Errorf("Expectd wetland to have 1 card, got %d", board.GetHabitat(Wetland).CountCards())
    }

    if board.GetHabitat(Grassland).CountCards() != 1 {
        t.Errorf("Expectd grassland to have 1 card, got %d", board.GetHabitat(Grassland).CountCards())
    }
}

func TestItDoesNotPlaceMoreThanTheNumberOfColumns(t *testing.T) {
    board := CreateBoard(
        CreateRow(Forest, 2, GainFood),
    )

    card := CreateCard("Cockatiel", 5, nil, 0, Forest)
    other := CreateCard("Mocking", 5, nil, 0, Forest)
    yetanother := CreateCard("Goose", 5, nil, 0, Forest)

    board.PlaceCard(card)
    board.PlaceCard(other)
    index := board.PlaceCard(yetanother)

    if index != -1 {
        t.Errorf("Expected card to not be placed, instead got position: %d", index)
    }
}
