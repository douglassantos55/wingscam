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
