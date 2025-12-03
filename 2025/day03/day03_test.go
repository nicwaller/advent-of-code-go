package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxJoltage(t *testing.T) {
	testCases := map[string]int{
		"987654321111111": 98,
		"811111111111119": 89,
		"234234234234278": 78,
		"818181911112111": 92,
	}
	for input, expected := range testCases {
		actual := maxJoltage(input)
		t.Log(input)
		assert.Equal(t, expected, actual)
	}
}
