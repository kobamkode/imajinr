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
		wr.NewChoice("walk", 6),
		wr.NewChoice("meet monster", 3),
		wr.NewChoice("woah item!", 1),
	)

	res := chooser.Pick()
	fmt.Println(res)
}
