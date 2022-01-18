package main

type Player struct {
	deck  *Deck
	foods []Food
	board [][]*Card
	hand  *LinkedList
}

func CreatePlayer(deck *Deck) *Player {
	return &Player{
		deck: deck,
		hand: &LinkedList{},
		board: [][]*Card{
			{nil, nil, nil, nil, nil},
			{nil, nil, nil, nil, nil},
			{nil, nil, nil, nil, nil},
		},
	}
}

func (player *Player) CountCardsInHand() int {
	return player.hand.Count()
}

func (player *Player) DrawCard(card *Card) {
	if card != nil {
		player.hand.Push(card)
	}
}

func (player *Player) PlayCard(card *Card) {
    for i := 0; i < 3; i++ {
        for j := 0; j < 5; j++ {
            if player.board[i][j] == nil {
                player.board[i][j] = card
            }
        }
    }
}

func (player *Player) DrawCards(qty int) {
	for i := 0; i < qty; i++ {
		card := player.deck.Draw()
		player.hand.Push(card)
	}
}

func (player *Player) GainFood(food Food) {
	player.foods = append(player.foods, food)
}

func (player *Player) HasFood(food Food) bool {
	for _, foodType := range player.foods {
		if food == foodType {
			return true
		}
	}
	return false
}

func (player *Player) PayFood(condition Condition) bool {
	result := condition.Apply(player.foods)

	if len(result) != 0 {
		for i := range result {
			player.foods = append(player.foods[:i], player.foods[i+1:]...)
		}

		return true
	}

	return false
}

func (player *Player) PayEggs(qty uint8) bool {
	sum := uint8(0)
	candidates := []*Card{}

	for i := 0; i < 3; i++ {
		for j := 0; j < 5; j++ {
			card := player.board[i][j]

			if card != nil {
				if card.CountEggs() > 0 {
					sum += card.CountEggs()
					candidates = append(candidates, card)

					if sum >= qty {
						for _, candidate := range candidates {
							qty -= candidate.DropEggs(qty)
						}

						return true
					}
				}
			}
		}
	}

	return false
}
