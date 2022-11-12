package grid

import (
	"fmt"
	"testing"
)

func TestOffsetSimple(t *testing.T) {
	g := NewGrid[int](2, 3)
	expectCell := func(expected int, loc []int) {
		actual := g.OffsetFromCell(loc)
		expect(t, expected, actual, fmt.Sprintf("Offset for cell %v", loc))
	}
	expectCell(0, []int{0, 0})
	expectCell(1, []int{0, 1})
	expectCell(2, []int{0, 2})
	expectCell(3, []int{1, 0})
	expectCell(4, []int{1, 1})
	expectCell(5, []int{1, 2})
}

func TestOffsetComplex(t *testing.T) {
	g := NewGrid[int](2, 3)
	g.offsets = []int{10, 10}
	expectCell := func(expected int, loc []int) {
		actual := g.OffsetFromCell(loc)
		expect(t, expected, actual, fmt.Sprintf("Offset for cell %v", loc))
	}
	expectCell(0, []int{10, 10})
	expectCell(1, []int{10, 11})
	expectCell(2, []int{10, 12})
	expectCell(3, []int{11, 10})
	expectCell(4, []int{11, 11})
	expectCell(5, []int{11, 12})
}

func TestGrid_IsInGrid(t *testing.T) {
	g := NewGrid[bool](3, 3)
	centroid := Cell{1, 1}
	expect[bool](t, true, g.IsInGrid(centroid), "centroid in grid")

	for _, n := range g.NeighboursAdjacent(centroid, true) {
		expect(t, true, g.IsInGrid(n), "neighbour in grid")
	}

	corner := Cell{0, 0}
	cornerNeighbours := g.NeighboursAdjacent(corner, false)
	expect(t, 2, len(cornerNeighbours), "number of neighbours for corner cell")
}

func expect[T comparable](t *testing.T, expected T, actual T, name string) {
	if expected != actual {
		t.Errorf("%s:\nExpected: %v\nActual: %v\n", name, expected, actual)
	}
}
