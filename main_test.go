package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestImajinr(t *testing.T) {
	// t.Run("", func(t *testing.T) {})
	t.Run("create new character", func(t *testing.T) {
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
	})
}
