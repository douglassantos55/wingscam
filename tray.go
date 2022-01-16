package main

type Tray struct {
	cards []*Card
	size  uint8
}

func CreateTray(size uint8) *Tray {
	return &Tray{cards: []*Card{}, size: 3}
}

func (tray *Tray) Count() uint8 {
	return uint8(len(tray.cards))
}

func (tray *Tray) Replenish(deck *Deck) {
	tray.cards = []*Card{}

	for i := uint8(0); i < tray.size; i++ {
		tray.cards = append(tray.cards, deck.Draw())
	}
}

func (tray *Tray) GetCard(index uint8) *Card {
    if index < tray.size {
        card := tray.cards[index]
        tray.cards = append(tray.cards[:index], tray.cards[index+1:]...)

        return card
    }

    return nil
}
