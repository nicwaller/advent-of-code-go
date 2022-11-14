package grid

import "advent-of-code/lib/iter"

type Grid[T comparable] struct {
	storage []T
	// Grid invariant: len(offsets) == len(dimensions)
	dimensions []int // ordered major to minor
	offsets    []int // ordered major to minor
	jumps      []int // pre-calculated cache of how much to jump for each dimension
}

type Cell []int

func IsZeroCell(c Cell) bool {
	for _, v := range c {
		if v != 0 {
			return false
		}
	}
	return true
}

func (grid *Grid[T]) recalculateJumps() {
	grid.jumps = make([]int, len(grid.dimensions))
	jumpSize := 1
	for d := len(grid.dimensions) - 1; d >= 0; d-- {
		grid.jumps[d] = jumpSize
		jumpSize *= grid.dimensions[d]
	}
}

// OffsetFromCell returns the index of the storage slice that corresponds to the requested cell
// 2D example, 3x3 grid:
//   0 1 2
//   3 4 5
//   6 7 8
func (grid *Grid[T]) OffsetFromCell(cell Cell) int {
	pos := 0
	for d, z := range cell {
		pos += (z - grid.offsets[d]) * grid.jumps[d]
	}
	return pos
}

func (grid *Grid[T]) CellFromOffset(offset int) Cell {
	c := make([]int, len(grid.dimensions))
	for d, j := range grid.jumps {
		c[d] = offset / j
		offset -= c[d] * j
	}
	for d, ofs := range grid.offsets {
		c[d] += ofs
	}
	return c
}

func (grid *Grid[T]) Get(cell Cell) T {
	return grid.storage[grid.OffsetFromCell(cell)]
}

func (grid *Grid[T]) Set(cell Cell, v T) {
	grid.storage[grid.OffsetFromCell(cell)] = v
}

func (grid *Grid[T]) Values() *[]T {
	return &grid.storage
}

func (grid *Grid[T]) ValuesIterator() iter.Iterator[T] {
	return iter.ListIterator(grid.storage)
}

func (grid *Grid[T]) All() Slice {
	s := make([]Range, len(grid.dimensions))
	for d := 0; d < len(grid.dimensions); d++ {
		s[d] = Range{
			Origin:   grid.offsets[d],
			Terminus: grid.offsets[d] + grid.dimensions[d],
		}
	}
	return s
}

func (grid *Grid[T]) NeighboursAdjacent(c Cell, includeCentre bool) []Cell {
	possibilities := make([]Cell, 0)
	// this is hard to write n-dimensionally!
	nDimensions := len(grid.dimensions)
	switch nDimensions {
	case 0:
		panic("grid cannot have 0 dimensions")
	case 1:
		possibilities = []Cell{
			[]int{c[0] - 1},
			[]int{c[0] + 1},
		}
	case 2:
		possibilities = []Cell{
			[]int{c[0] - 1, c[1]},
			[]int{c[0] + 1, c[1]},
			[]int{c[0], c[1] - 1},
			[]int{c[0], c[1] + 1},
		}
	case 3:
		possibilities = []Cell{
			[]int{c[0], c[1], c[2] - 1},
			[]int{c[0], c[1], c[2] + 1},
			[]int{c[0], c[1] - 1, c[2]},
			[]int{c[0], c[1] + 1, c[2]},
			[]int{c[0] - 1, c[1], c[2]},
			[]int{c[0] + 1, c[1], c[2]},
		}
	default:
		panic("neighbours not implemented for higher dimensions")
	}
	if includeCentre {
		possibilities = append(possibilities, c)
	}
	return iter.ListIterator[Cell](possibilities).Filter(grid.IsInGrid).List()
}

func (grid *Grid[T]) NeighboursSurround(c Cell, includeCentre bool) []Cell {
	possibilities := make([]Cell, 0)
	nDimensions := len(grid.dimensions)
	switch nDimensions {
	case 0:
		panic("grid cannot have 0 dimensions")
	case 1:
		possibilities = []Cell{
			[]int{c[0] - 1},
			[]int{c[0] + 1},
		}
	case 2:
		possibilities = []Cell{
			[]int{c[0] - 1, c[1] - 1},
			[]int{c[0] - 1, c[1] + 0},
			[]int{c[0] - 1, c[1] + 1},
			[]int{c[0] + 0, c[1] - 1},
			//[]int{c[0] + 0, c[1] + 0},
			[]int{c[0] + 0, c[1] + 1},
			[]int{c[0] + 1, c[1] - 1},
			[]int{c[0] + 1, c[1] + 0},
			[]int{c[0] + 1, c[1] + 1},
		}
	default:
		panic("neighbours not implemented for higher dimensions")
	}
	if includeCentre {
		possibilities = append(possibilities, c)
	}
	return iter.ListIterator[Cell](possibilities).Filter(grid.IsInGrid).List()
}

func (grid *Grid[T]) IsInGrid(c Cell) bool {
	for dim, pos := range c {
		if pos < grid.offsets[dim] {
			return false
		}
		if pos >= grid.offsets[dim]+grid.dimensions[dim] {
			return false
		}
	}
	return true
}
