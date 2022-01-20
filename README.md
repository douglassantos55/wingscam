// Game creation
cards := map[string]*Card = {
    // all the cards in the game
}

// Create deck with all the cards
deck := CreateDeck(cards...)

tray := CreateTray(3)
tray.Replenish(deck)

player := CreatePlayer(deck, CreateBoard())
anotherPlayer := CreatePlayer(deck, CreateBoard())

player::Play_Card{
    card: "089i8aihtotb1",
    habitat: nil or (0 -> Forest, 1 -> Grassland, 2 -> Wetland),
}
card := cards[event.card]

if !player.Pay(card.foodCost) {
    Send("not enough food")
}

row := player.Board.GetRow(event.Habitat || card.Habitat)

if !player.PayEggs(column.EggsCost) {
    Send("not enough eggs")
}

row.Place(card) // column.Place() would check if the card has a "WhenPlayed" power and activate it
Send("end_turn") // playing card does not activate the row's birds power
