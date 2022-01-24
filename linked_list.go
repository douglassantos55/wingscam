package main

type LinkedList struct {
	head *Node
}

func (list *LinkedList) Push(card *Card) {
	if card != nil {
		node := &Node{card: card, next: nil}

		if list.head == nil {
			list.head = node
		} else {
			var last *Node = nil
			for cur := list.head; cur != nil; cur = cur.next {
				last = cur
			}
			last.next = node
		}
	}
}

func (list *LinkedList) Count() int {
	count := 0
	for cur := list.head; cur != nil; cur = cur.next {
		count++
	}
	return count
}

func (list *LinkedList) Get(index uint8) *Card {
	cur := list.head
	var prev *Node = nil

	for i := uint8(0); i < index; i++ {
		if cur == nil || cur.next == nil {
			return nil
		}

		prev = cur
		cur = cur.next
	}

	if prev != nil {
		prev.next = cur.next
	} else {
		list.head = cur.next
	}

	return cur.card
}

func (list *LinkedList) All() []*Card {
    cards := []*Card{}

	for cur := list.head; cur != nil; cur = cur.next {
        cards = append(cards, cur.card)
	}

	return cards
}
