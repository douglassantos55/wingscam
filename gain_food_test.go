package main

import (
	"testing"
)

func TestFeederIsCreatedWithRandomFoods(t *testing.T) {
	// Use higher quantities to reduce change of failing
	feeder1 := CreateFeeder(10)
	feeder2 := CreateFeeder(10)

	equals := uint8(0)

	for i, food := range feeder1.foods {
		if feeder2.foods[i] == food {
			equals++
		}
	}

	if equals == feeder1.Count() {
		t.Errorf("Feeder should have random Food")
	}
}

func TestFoodIsRemovedFromFeeder(t *testing.T) {
	feeder := CreateFeeder(5)
	feeder.GetFood(2)

	if feeder.Count() != 4 {
		t.Errorf("Should have removed food from feeder")
	}
}

func TestPlayerGetsFood(t *testing.T) {
	feeder := CreateFeeder(5)
	player := CreatePlayer(CreateDeck())

	food := feeder.GetFood(1)
	player.GainFood(food)

	if !player.HasFood(food) {
		t.Errorf("Player should have a Fish")
	}
}

func TestFeederIsRefilledWhenItHasOnlyOneFood(t *testing.T) {
	feeder := CreateFeeder(1)
	feeder.GetFood(0)

	if feeder.Count() != 4 {
		t.Errorf("Feeder should have been refilled")
	}
}

func TestFeederIsRefilledWhenItHasNoFood(t *testing.T) {
	feeder := CreateFeeder(0)
	feeder.GetFood(0)

	if feeder.Count() != 4 {
		t.Errorf("Feeder should have been refilled")
	}
}

func TestFeederIsRefilledWhenItHasAllSameFood(t *testing.T) {
	feeder := CreateFeeder(5)
	feeder.foods = []Food{Fish, Fish, Fish, Fish, Fish}
	feeder.GetFood(4)

	if feeder.Count() != 4 {
		t.Errorf("Feeder should have been refilled")
	}
}
