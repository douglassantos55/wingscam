package main

type Action interface{
    Execute()
}

type Power struct{
    trigger Trigger
    action Action
}

type LayEggAction struct {
	qty uint8
    card *Card
}

func (action LayEggAction) Execute() {
    action.card.LayEggs(action.qty)
}

type Trigger uint8

const (
	WhenPlayed    Trigger = iota
	WhenActivated         = iota
)
