package main

type Food uint8

func (food Food) String() string {
    names := map[Food]string {
        Seed: "Seed",
        Fish: "Fish",
        Rodent: "Rodent",
        Fruit: "Fruit",
        Invertebrate: "Invertebrate",
    }

    return names[food]
}

const (
	Seed         Food = iota
	Fish              = iota
	Rodent            = iota
	Fruit             = iota
	Invertebrate      = iota
)
