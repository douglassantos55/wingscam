package main

type Card struct {
	name      string
	eggs      uint8
	eggsLimit uint8
	foodCost  Condition
}

func CreateCard(name string, eggsLimit uint8, foodCost Condition) *Card {
	return &Card{
		name:      name,
		eggsLimit: eggsLimit,
		foodCost:  foodCost,
	}
}

func (card *Card) CountEggs() uint8 {
	return card.eggs
}

func (card *Card) LayEggs(qty uint8) {
	for i := uint8(0); i < qty; i++ {
		if card.eggs < card.eggsLimit {
			card.eggs++
		}
	}
}
