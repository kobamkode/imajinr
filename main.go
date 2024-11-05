package main

import (
	"time"
)

func main() {
	// TODO: check if character is exist, if not create new
	c := NewCharacter("john")
	for {
		time.Sleep(1 * time.Second)
		action(c)
	}
}
