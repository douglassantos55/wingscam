package main

type Board struct {
	rows []*Row
}

func CreateBoard(rows ...*Row) *Board {
	return &Board{
		rows: rows,
	}
}

func (board *Board) GetHabitat(habitat Habitat) *Row {
	for _, row := range board.rows {
		if row.GetHabitat() == habitat {
			return row
		}
	}

	return nil
}

func (board *Board) PlaceCard(card *Card) int {
	row := board.GetHabitat(card.habitat)

	if row != nil {
		return row.PlaceCard(card)
	}

    return -1
}

func (board *Board) GetCards() []*Card {
    cards := []*Card{}

    for _, row  := range board.rows {
        cards = append(cards, row.GetCards()...)
    }

    return cards
}

type Row struct {
	habitat Habitat
	columns []*Column
}

func CreateRow(habitat Habitat, columns uint8) *Row {
	cols := []*Column{}

	for i := uint8(0); i < columns; i++ {
		cols = append(cols, &Column{})
	}

	return &Row{
		habitat: habitat,
		columns: cols,
	}
}

func (row *Row) GetHabitat() Habitat {
	return row.habitat
}

func (row *Row) GetCards() []*Card {
    cards := []*Card{}

	for _, col := range row.columns {
        if !col.Empty() {
            cards = append(cards, col.GetCard())
        }
    }

    return cards
}

func (row *Row) CountCards() uint8 {
    return uint8(len(row.GetCards()))
}

func (row *Row) PlaceCard(card *Card) int {
	for i, col := range row.columns {
		if col.Empty() {
			col.PlaceCard(card)
            row.ActivatePowers(i)
            return i
		}
	}

    return -1
}

func (row *Row) ActivatePowers(start int) {
    for i := start-1; i >= 0; i-- {
        card := row.columns[i].card

        if card != nil {
            card.Trigger(WhenActivated)
        }
    }
}

type Column struct {
	card *Card
}

// int(i / 2) + 2 --> eggs
// int(i / 2) + 1 --> food and card
func (col *Column) Empty() bool {
	return col.card == nil
}

func (col *Column) PlaceCard(card *Card) {
	col.card = card
    card.Trigger(WhenPlayed)
}

func (col *Column) GetCard() *Card {
    return col.card
}
