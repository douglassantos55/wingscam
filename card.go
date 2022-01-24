package main

import (
	"encoding/json"
	"fmt"
)

type Card struct {
	Name      string
	eggs      uint8
	EggsLimit uint8
	eggsCost  uint8
	Habitat   Habitat
	FoodCost  Condition
	powers    map[Trigger]Power
}

func (card *Card) UnmarshalJSON(data []byte) error {
	var v map[string]interface{}

	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	fmt.Println(v)

	habitatMap := map[string]Habitat{
		"forest":    Forest,
		"grassland": Grassland,
		"wetland":   Wetland,
	}

    foodMap := map[string]Food {
        "fruit": Fruit,
        "fish": Fish,
        "rodent": Rodent,
        "invertebrate": Invertebrate,
    }

	costMap := map[string]interface{}{
		"and": And,
		"or":  Or,
        "single": And,
	}

	costType := v["costType"].(string)
	costReqs := []Food{foodMap[v["foodCost"].(string)]}
	costFunc := costMap[costType].(func(...Food) Condition)


	card.Name = v["name"].(string)
	card.EggsLimit = uint8(v["eggsLimit"].(float64))
	card.Habitat = habitatMap[v["habitat"].(string)]
	card.FoodCost = costFunc(costReqs...)

	return nil
}

func CreateCard(name string, eggsLimit uint8, foodCost Condition, eggsCost uint8, habitat Habitat) *Card {
	return &Card{
		Name:      name,
		EggsLimit: eggsLimit,
		eggsCost:  eggsCost,
		FoodCost:  foodCost,
		Habitat:   habitat,
		powers:    make(map[Trigger]Power),
	}
}

func (card *Card) CountEggs() uint8 {
	return card.eggs
}

func (card *Card) LayEggs(qty uint8) {
	for i := uint8(0); i < qty; i++ {
		if card.eggs < card.EggsLimit {
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
	return card.FoodCost
}

func (card *Card) GetEggsCost() uint8 {
	return card.eggsCost
}

func (card *Card) GivePower(trigger Trigger, action Power) {
	card.powers[trigger] = action
}

func (card *Card) Trigger(trigger Trigger) {
	power := card.powers[trigger]

	if power != nil {
		power.Execute()
	}
}

func (card *Card) String() string {
    return fmt.Sprintf("%s (eggs: %d)", card.Name, card.CountEggs())
}
