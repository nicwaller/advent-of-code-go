package main

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"advent-of-code/lib/f8l"
	"advent-of-code/lib/iter"
	"advent-of-code/lib/util"
)

func TestMaxJoltage(t *testing.T) {
	type TestCase struct {
		bank    string
		joltage int
	}
	tests := []TestCase{
		{bank: "999999999999000", joltage: 999999999999},
		{bank: "987654321111111", joltage: 987654321111},
		{bank: "811111111111119", joltage: 811111111119},
		{bank: "234234234234278", joltage: 434234234278},
		{bank: "818181911112111", joltage: 888911112111},
	}
	for _, tc := range tests {
		t.Run(tc.bank, func(t *testing.T) {
			bank := f8l.Map(iter.StringIterator(tc.bank, 1).List(), util.UnsafeAtoi)
			actual := intify(maxJoltageLong(bank, 12))
			t.Log(tc.bank, actual)
			assert.Equal(t, tc.joltage, actual)
		})
	}
}

func TestIntify(t *testing.T) {
	assert.Equal(t, 123, intify([]int{1, 2, 3}))
}
