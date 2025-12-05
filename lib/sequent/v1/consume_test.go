package v1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTakeOne(t *testing.T) {
	ite := Range(2, 3)
	assert.Equal(t, 2, ite.TakeOne())
	assert.Panics(t, func() {
		ite.TakeOne()
	})
}

func TestTakeN(t *testing.T) {
	ite := Range(2, 7)
	assert.Equal(t, []int{}, ite.TakeN(0))
	assert.Equal(t, []int{2}, ite.TakeN(1))
	assert.Equal(t, []int{3, 4}, ite.TakeN(2))
	assert.Panics(t, func() {
		ite.TakeN(3)
	})
	assert.Zero(t, ite.Count())
}
