package main

type Condition interface {
	Apply(source []Food) []uint8
}

type AndCondition struct {
	requirements []Food
}

type OrCondition struct {
	requirements []Food
}

type SingleCondition struct {
    requirement Food
}

func (cond SingleCondition) Apply(source []Food) []uint8 {
	for index, food := range source {
        if cond.requirement == food {
            return []uint8{uint8(index)}
        }
	}

    return []uint8{}
}

func (cond OrCondition) Apply(source []Food) []uint8 {
	matches := []uint8{}

	for index, food := range source {
		for _, req := range cond.requirements {
			if req == food {
				matches = append(matches, uint8(index))
			}
		}
	}

	return matches
}

func (cond AndCondition) Apply(source []Food) []uint8 {
	matches := map[Food]uint8{}

	for index, food := range source {
		for _, req := range cond.requirements {
			if req == food {
				matches[req] = uint8(index)
			}
		}
	}

	if len(matches) != len(cond.requirements) {
		return []uint8{}
	}

	v := make([]uint8, 0, len(matches))
	for _, value := range matches {
		v = append(v, value)
	}

	return v
}

func And(foods ...Food) AndCondition {
	return AndCondition{requirements: foods}
}

func Or(foods ...Food) OrCondition {
	return OrCondition{requirements: foods}
}

func Single(food Food) SingleCondition {
	return SingleCondition{requirement: food}
}
