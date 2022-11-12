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

func (grid *Grid[T]) Filter(filterFn func(v T) bool) iter.Iterator[Cell] {
	return grid.Cells().Filter(func(cell Cell) bool {
		offset := grid.OffsetFromCell(cell)
		value := grid.storage[offset]
		return filterFn(value)
	})
}
