package v1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIterator_Map(t *testing.T) {
	add1 := func(v int) int { return v + 1 }
	original := Range(1, 4)
	result := original.Map(add1).List()
	assert.Equal(t, []int{2, 3, 4}, result)
}

func TestIterator_Reduce(t *testing.T) {
	adder := func(i, j int) int { return i + j }
	assert.Equal(t, 6, Range(1, 4).Reduce(0, adder))
}

func TestChain(t *testing.T) {
	ite := Chain(
		Range(0, 2),
		Range(2, 4),
	)
	assert.Equal(t, []int{0, 1, 2, 3}, ite.List())
}

func TestIterator_Counting(t *testing.T) {
	counter := 0
	Range(0, 4).Counting(&counter).Each(nil)
	assert.Equal(t, 4, counter)
}

func TestChunk(t *testing.T) {
	t.Run("whole chunks", func(t *testing.T) {
		assert.Equal(t, [][]int{
			{0, 1},
			{2, 3},
			{4, 5},
		}, Chunk(Range(0, 6), 2).List())
	})

	t.Run("partial chunk", func(t *testing.T) {
		assert.Equal(t, [][]int{
			{0, 1},
			{2},
		}, Chunk(Range(0, 3), 2).List())
	})

	t.Run("empty", func(t *testing.T) {
		assert.Equal(t, [][]int{}, Chunk(Range(0, 0), 2).List())
	})

}

func TestSlidingWindow(t *testing.T) {
	ite := Range(0, 4)
	iteSlider := SlidingWindow(ite, 2)
	assert.Equal(t, [][]int{
		{0, 1},
		{1, 2},
		{2, 3},
	}, iteSlider.List())
}

func TestEnumerate(t *testing.T) {
	ite := StringIterator("hi")
	expected := []IndexedValue[string]{
		{Index: 0, Value: "h"},
		{Index: 1, Value: "i"},
	}
	assert.Equal(t, expected, Enumerate(ite).List())
	assert.Zero(t, ite.Count())
}
