package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCharacter(t *testing.T) {
	name := "john"
	got := NewCharacter(name)
	want := &Character{
		name: name,
		job:  "novice",
		hp:   100,
		mp:   100,
		atk:  10,
	}
	assert.Equal(t, want, got)
}
