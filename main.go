package main

import (
	"fmt"
	"time"

	wr "github.com/mroth/weightedrand/v2"
)

func main() {
	for {
		time.Sleep(1 * time.Second)
		action()
	}
}

func action() {
	chooser, _ := wr.NewChooser(
		wr.NewChoice("walk", 7),
		wr.NewChoice("fight", 2),
		wr.NewChoice("loot", 1),
	)

	res := chooser.Pick()
	fmt.Println(res)
}
