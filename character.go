package main

import "errors"

type Character struct {
	name string
}

func New() *Character {
	return &Character{}
}

func (c *Character) SetName(name string) error {
	if name != "" {
		c.name = name
		return nil
	}

	return errors.New("name should not empty")
}

func (c *Character) Name() string {
	return c.name
}
