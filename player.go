package main

type Player struct {
	deck  *Deck
	foods []Food
	hand  *LinkedList
}

func CreatePlayer(deck *Deck) *Player {
	return &Player{deck: deck, hand: &LinkedList{}}
}

func (player *Player) CountCardsInHand() int {
	return player.hand.Count()
}

func (player *Player) DrawCard(card *Card) {
    if card != nil {
        player.hand.Push(card)
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
