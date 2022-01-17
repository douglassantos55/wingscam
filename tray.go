package main

type Tray struct {
	cards *LinkedList
	size  uint8
}

func CreateTray(size uint8) *Tray {
	return &Tray{cards: &LinkedList{}, size: 3}
}

func (tray *Tray) Count() uint8 {
	return uint8(tray.cards.Count())
}

func (tray *Tray) Replenish(deck *Deck) {
	tray.cards = &LinkedList{}

	for i := uint8(0); i < tray.size; i++ {
		tray.cards.Push(deck.Draw())
	}
}

func (tray *Tray) GetCard(index uint8) *Card {
	if index < tray.size {
		return tray.cards.Get(index)
	}

	return nil
}
