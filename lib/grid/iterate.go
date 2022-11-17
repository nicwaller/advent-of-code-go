package grid

import (
	"advent-of-code/lib/iter"
)

func (grid *Grid[T]) Cells() iter.Iterator[Cell] {
	offset := -1
	var curCell Cell
	return iter.Iterator[Cell]{
		Next: func() bool {
			offset++
			if offset < 0 {
				panic("offset must never be negative") // DEBUG
			}
			if offset > len(grid.storage) {
				panic("offset is too high")
			}
			curCell = grid.CellFromOffset(offset)
			return offset < len(grid.storage)
		},
		Value: func() Cell {
			return curCell
		},
	}
}

func (grid *Grid[T]) Filter(filterFn func(Cell, T) bool) iter.Iterator[Cell] {
	return grid.Cells().Filter(func(cell Cell) bool {
		offset := grid.OffsetFromCell(cell)
		value := grid.storage[offset]
		return filterFn(cell, value)
	})
}

func (grid *Grid[T]) FloodSelect(origin Cell, selFn func(v T) bool) []Cell {
	// I would rather use sets/maps here, but Cells don't implement comparable
	// so the implementation becomes quite ugly
	seen := make([]Cell, 0)
	grid.floodSelect(origin, selFn, &seen)
	return seen
}

func (grid *Grid[T]) floodSelect(current Cell, selFn func(v T) bool, seen *[]Cell) {
	// PERF: yes, I know this is quite inefficient :(
	// but it seems that basins are never larger than 50 cells
	// it's not too bad with small N.
	// perhaps a more efficient option would be to allocate a parallel Grid of booleans?
	wasSeen := func(cell Cell) bool {
		for _, s := range *seen {
			match := true
			for d, _ := range s {
				if cell[d] != s[d] {
					match = false
				}
			}
			if match {
				return true
			}
		}
		return false
	}
	for _, neighbour := range grid.NeighboursAdjacent(current, false) {
		// run selFn first because it'll be cheaper for this problem (Day 09)
		if !selFn(grid.Get(neighbour)) {
			continue
		}
		if wasSeen(neighbour) {
			continue
		}
		*seen = append(*seen, neighbour)
		grid.floodSelect(neighbour, selFn, seen)
	}
}
