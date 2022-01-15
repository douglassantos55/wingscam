package main

import (
	"crypto/rand"
	"math/big"
)

type Feeder struct {
	foods []Food
	count uint8
}

func CreateFeeder(initial uint8) *Feeder {
	feeder := &Feeder{}
	feeder.fill(initial)

	return feeder
}

func (feeder *Feeder) fill(initial uint8) {
	for i := uint8(0); i < initial; i++ {
		random, _ := rand.Int(rand.Reader, big.NewInt(Fish+1))
		feeder.foods = append(feeder.foods, Food(random.Uint64()))
	}
}

func (feeder *Feeder) refill() {
	feeder.fill(5 - feeder.Count())
}

func (feeder *Feeder) GetFood(foodType Food) Food {
	if feeder.Count() <= 1 {
		feeder.refill()
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
