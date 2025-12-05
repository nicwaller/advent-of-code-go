package v1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIterator_Each(t *testing.T) {
	callbackCounter := 0
	Range(0, 4).Each(func(v int) {
		callbackCounter++
	})
	assert.Equal(t, 4, callbackCounter)

	assert.NotPanics(t, func() {
		Range(0, 4).Each(nil)
	})
}

func TestIterator_Count(t *testing.T) {
	assert.Equal(t, 4, Range(0, 4).Count())
}

func TestSum(t *testing.T) {
	assert.Equal(t, 10, Sum(Range(1, 5)))
}
