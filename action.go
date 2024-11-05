package main

import (
	wr "github.com/mroth/weightedrand/v2"
)

func action(c *Character) {
	chooser, _ := wr.NewChooser(
		wr.NewChoice("walk", 60),
		wr.NewChoice("fight", 30),
		wr.NewChoice("loot", 10),
	)

	switch chooser.Pick() {
	case "walk":
		walking()
	case "fight":
		fighting()
	case "loot":
		looting()
	}
}

func looting() {
	panic("unimplemented")
}

func fighting() {
	panic("unimplemented")
}

func walking() {
	panic("unimplemented")
}
