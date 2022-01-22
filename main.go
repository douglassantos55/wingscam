package main

func main() {
	cards := []*Card{
		CreateCard("", 0, nil, 0, Forest),
	}

	deck := CreateDeck(cards...)
	game := CreateGame(CreatePlayer(deck), CreatePlayer(deck))
}
