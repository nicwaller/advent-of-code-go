package v1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPermute(t *testing.T) {
	count := 0
	p := Permute([]string{"r", "g", "b"}).Counting(&count)
	assert.Equal(t, [][]string{
		{"r", "g", "b"},
		{"g", "r", "b"},
		{"b", "r", "g"},
		{"r", "b", "g"},
		{"g", "b", "r"},
		{"b", "g", "r"},
	}, p.List())
	assert.Equal(t, 6, count)
	assert.Zero(t, p.Count())
}
