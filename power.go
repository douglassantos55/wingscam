package main

type Power interface{
    Execute()
}

type LayEggPower struct {
	qty uint8
    card *Card
}

func (power LayEggPower) Execute() {
    power.card.LayEggs(power.qty)
}

type Trigger uint8

const (
	WhenPlayed    Trigger = iota
	WhenActivated         = iota
)
