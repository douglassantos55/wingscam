package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
)

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
	chooser Chooser
}

func CreateGame(chooser Chooser, players ...*Player) *Game {
	return &Game{
		chooser: chooser,
		current: players[0],
		players: players,
		feeder:  CreateFeeder(5),
	}
}

func (game *Game) Start() {
	for _, player := range game.players {
		player.DrawCards(5)
		player.GainFood(Seed)
		player.GainFood(Fish)
		player.GainFood(Fruit)
		player.GainFood(Rodent)
		player.GainFood(Invertebrate)
	}

	for {
		fmt.Println("Choose an action: ")
		fmt.Println("1: Play card")
		fmt.Println("2: Gain food")
		fmt.Println("3: Lay eggs")
		fmt.Println("4: Draw cards")
		fmt.Println("5: See your hand")
		fmt.Println("6: See your board")
		fmt.Println("7: See your food")

		switch ReadInput() {
		case "1":
			game.PlayCard()
		case "2":
			game.GainFood()
		case "3":
			game.LayEggs()
		case "4":
			game.DrawCards()
		case "5":
			for cur := game.current.hand.head; cur != nil; cur = cur.next {
				fmt.Printf("%s\n", cur.card.Name)
			}

			fmt.Print("\n")
		case "6":
			for _, row := range game.current.board.rows {
				fmt.Print("Habitat: ", row.GetHabitat())

				for _, card := range row.GetCards() {
					fmt.Printf("\t %s (%d)", card.Name, card.CountEggs())
				}

				fmt.Print("\n")
			}
		case "7":
			for i, food := range game.current.foods {
				fmt.Printf("(%d) %s\n", i, food)
			}
		}
	}
}

type Chooser interface {
	Choose(items interface{}) int
}
type TerminalChooser struct{}

func (t TerminalChooser) Choose(items interface{}) int {
	v := reflect.ValueOf(items)
	stringers := make([]fmt.Stringer, v.Len())

	for i := 0; i < v.Len(); i++ {
		stringers[i] = v.Index(i)
	}

	for i, item := range stringers {
		fmt.Printf("(%d) %s\n", i, item)
	}

	fmt.Print("\n Choose one: ")

	index, err := strconv.ParseUint(ReadInput(), 10, 64)

	if err != nil {
		return -1
	}

	return int(index)
}

func ReadInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	return scanner.Text()
}

func (game *Game) GainFood() {
	row := game.current.board.GetActionRow(GainFood)

	if row != nil {
		qty := row.CountCards()/2 + 1

		for i := uint8(0); i < qty; i++ {
			index := game.chooser.Choose(game.feeder.foods)

			if index != -1 {
				game.current.GainFood(game.feeder.GetFood(uint8(index)))
			}
		}

		row.ActivatePowers()
	}
}

func (game *Game) PlayCard() {
	index := game.chooser.Choose(game.current.hand.All())

	if index != -1 {
		card := game.current.hand.Get(uint8(index))
		if card.GetFoodCost() != nil {
			game.current.PayFood(card.GetFoodCost())
		}
		game.current.PlayCard(card)
	}
}

func (game *Game) LayEggs() {
	cards := game.current.board.GetCards()
	row := game.current.board.GetActionRow(LayEggs)

	if row != nil {
		qty := row.CountCards()/2 + 2
		index := game.chooser.Choose(cards)

		if index != -1 {
			cards[index].LayEggs(qty)
			row.ActivatePowers()
		}
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
