package main

import (
	"crypto/rand"
	"math/big"
)

type Feeder struct {
	foods []Food
}

func CreateFeeder(initial uint8) *Feeder {
	feeder := &Feeder{}
	feeder.fill(initial)

	return feeder
}

func (feeder *Feeder) fill(initial uint8) {
	feeder.foods = []Food{}

	for i := uint8(0); i < initial; i++ {
		random, _ := rand.Int(rand.Reader, big.NewInt(Fish+1))
		feeder.foods = append(feeder.foods, Food(random.Uint64()))
	}
}

func (feeder *Feeder) AllSame() bool {
	for i := uint8(0); i < feeder.Count()-2; i++ {
		if feeder.foods[i] != feeder.foods[i+1] {
			return false
		}
	}
	return true
}

func (feeder *Feeder) GetFood(foodType Food) Food {
	if feeder.Count() <= 1 || feeder.AllSame() {
		feeder.fill(5)
	}

	for i, food := range feeder.foods {
		if food == foodType {
			feeder.foods = append(feeder.foods[:i], feeder.foods[i+1:]...)
			return food
		}
	}

	panic("Food not in Feeder")
}

func (feeder *Feeder) Count() uint8 {
	return uint8(len(feeder.foods))
}
