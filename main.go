package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)
type Cards struct {
    Cards []*Card `json:"cards"`
}

func main() {
	file, err := os.Open("cards.json")

	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()
	byteValue, _ := ioutil.ReadAll(file)

    var cards Cards
	json.Unmarshal(byteValue, &cards)

	deck := CreateDeck(cards.Cards...)
	game := CreateGame(TerminalChooser{}, CreatePlayer(deck), CreatePlayer(deck))

	game.Start()
}
