package main

import (
	"log"
	"time"

	wr "github.com/mroth/weightedrand/v2"
)

// TODO: Get new character name from input field.
func GetNewCharacterName() string {
	return "John"
}

// TODO: Gets saved game from database.
func GetSavedGame() interface{} {
	return nil
}

type Character struct {
	name string
	job  string
	hp   int
	mp   int
	atk  int
}

func NewCharacter(name string) *Character {
	return &Character{
		name: name,
		job:  "novice",
		hp:   100,
		mp:   100,
		atk:  10,
	}
}

func (c *Character) Fight() {
	panic("unimplemented")
}

// Chooses by weighted random algorithm.
func ChooseAction() (string, error) {
	choose, err := wr.NewChooser(
		wr.NewChoice("walk", 100),
		wr.NewChoice("fight", 0),
	)
	if err != nil {
		return "", err
	}
	return choose.Pick(), nil
}

func Action(act string, c *Character) {
	if act == "walk" {
		log.Printf("%s is walking...", c.name)
	}

	if act == "fight" {
		c.Fight()
	}
}

func main() {
	game := GetSavedGame()
	char := new(Character)

	if game == nil {
		name := GetNewCharacterName()
		char = NewCharacter(name)
	}

	for {
		act, err := ChooseAction()
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(1 * time.Second)
		Action(act, char)
	}
}
