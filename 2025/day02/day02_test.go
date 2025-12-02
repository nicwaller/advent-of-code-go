package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsRepeatingSequence(t *testing.T) {
	t.Run("part-one", func(t *testing.T) {
		assert.True(t, isRepeatingSequenceN("55", 1))
		assert.True(t, isRepeatingSequenceN("6464", 2))
		assert.True(t, isRepeatingSequenceN("123123", 3))
	})

	t.Run("part-two", func(t *testing.T) {
		assert.True(t, isRepeatingSequence("12341234"))
		assert.True(t, isRepeatingSequenceN("12341234", 4))
		assert.True(t, isRepeatingSequence("123123123"))
		assert.True(t, isRepeatingSequenceN("123123123", 3))
		assert.True(t, isRepeatingSequence("1212121212"))
		assert.True(t, isRepeatingSequenceN("1212121212", 2))
		assert.True(t, isRepeatingSequence("1111111"))
		assert.True(t, isRepeatingSequenceN("1111111", 1))
	})
}
