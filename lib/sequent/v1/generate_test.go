package v1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRangeStepped(t *testing.T) {
	assert.Equal(t, []int{0, 2, 4}, RangeStepped(0, 5, 2).List())
	assert.Panics(t, func() {
		RangeStepped(0, 5, 0)
	})
}

func TestRuneIterator(t *testing.T) {
	assert.Equal(t, []rune{'h', 'i', '!'}, RuneIterator("hi!").List())
}
