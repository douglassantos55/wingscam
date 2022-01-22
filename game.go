package main

type Action uint8

const (
	GainFood  Action = iota
	LayEggs          = iota
	DrawCards        = iota
)

type Game struct {
	feeder  *Feeder
	current *Player
	players []*Player
}

func CreateGame(players ...*Player) *Game {
	return &Game{
		current: players[0],
		players: players,
		feeder:  CreateFeeder(5),
	}
}

func (game *Game) GainFood() {
	row := game.current.board.GetActionRow(GainFood)

	if row != nil {
		qty := row.CountCards()/2 + 1

		for i := uint8(0); i < qty; i++ {
			game.current.GainFood(game.feeder.GetFood(0))
		}

		row.ActivatePowers()
	}
}

func (game *Game) PlayCard() {
	card := game.current.hand.Get(0)
	game.current.PlayCard(card)
}

func (game *Game) LayEggs() {
	cards := game.current.board.GetCards()
	row := game.current.board.GetActionRow(LayEggs)

	if row != nil {
		qty := row.CountCards()/2 + 2
		cards[0].LayEggs(qty)

		row.ActivatePowers()
	}
}

func (game *Game) DrawCards() {
	row := game.current.board.GetActionRow(DrawCards)

	if row != nil {
		qty := row.CountCards()/2 + 1
		game.current.DrawCards(int(qty))

		row.ActivatePowers()
	}
}
