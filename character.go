package main

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
