package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCharacter(t *testing.T) {
	got := New()
	want := &Character{}
	assert.Equal(t, want, got)
}

func TestCharacter(t *testing.T) {

	t.Run("should has a name", func(t *testing.T) {
		c := Character{}
		err := c.SetName("john")
		want := Character{name: "john"}
		assert.NotNil(t, want)
		assert.NoError(t, err)
	})

	t.Run("should error when name not exist", func(t *testing.T) {
		c := Character{}
		err := c.SetName("")
		assert.Error(t, err)
	})

	t.Run("should return name", func(t *testing.T) {
		c := Character{}
		c.SetName("john")
		got := c.Name()
		want := "john"
		assert.Equal(t, want, got)
	})
}
