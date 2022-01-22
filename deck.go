package main

type Node struct {
	card *Card
	next *Node
}

func (deck *Deck) Pop() *Card {
	var card *Card = nil

	if deck.top != nil {
		card = deck.top.card
		deck.top = deck.top.next
	}

	return card
}

func (deck *Deck) Push(node *Node) {
	node.next = deck.top
	deck.top = node
}

type Deck struct {
	top *Node
}

func (deck *Deck) Draw() *Card {
	return deck.Pop()
}

func CreateDeck(cards ...*Card) *Deck {
	deck := new(Deck)
	for _, card := range cards {
		deck.Push(&Node{card: card, next: nil})
	}
	return deck
}
