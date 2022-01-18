package main

type Card struct {
	name      string
	eggs      uint8
	eggsLimit uint8
	eggsCost  uint8
	foodCost  Condition
}

func CreateCard(name string, eggsLimit uint8, foodCost Condition, eggsCost uint8) *Card {
	return &Card{
		name:      name,
		eggsLimit: eggsLimit,
		eggsCost:  eggsCost,
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

func (card *Card) DropEggs(qty uint8) uint8 {
    dropped := uint8(0)

	for i := uint8(0); i < qty; i++ {
		if card.eggs > 0 {
            dropped++
			card.eggs--
		}
	}

    return dropped
}

func (card *Card) GetFoodCost() Condition {
    return card.foodCost
}

func (card *Card) GetEggsCost() uint8 {
    return card.eggsCost
}
